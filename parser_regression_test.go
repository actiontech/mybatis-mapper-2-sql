package parser

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/actiontech/mybatis-mapper-2-sql/ast"
	"github.com/stretchr/testify/assert"
)

func TestParserSetWhereNoNullWHERE(t *testing.T) {
	actual, err := ParseXML(`
<mapper namespace="Test">
    <update id="updateToReverse">
        UPDATE t_sample_record t
        <set>t.cancelled_at = null</set>
        <where>
            <trim prefix="(" suffix=")">t.record_id=?</trim>
        </where>
    </update>
</mapper>`)
	if !assert.NoError(t, err) {
		return
	}
	assert.NotContains(t, actual, "nullWHERE")
	assert.Contains(t, actual, "NULL")
	assert.Contains(t, actual, "WHERE")
}

func TestParserWhereIfNoGlue(t *testing.T) {
	actual, err := ParseXML(`
<mapper namespace="Test">
    <select id="testWhereIfGlue">
        SELECT 1 FROM t
        <where>
            <if test="a != null">a=?</if>
            <if test="b != null">AND b=?</if>
        </where>
    </select>
</mapper>`)
	if !assert.NoError(t, err) {
		return
	}
	assert.NotContains(t, actual, "?AND")
	assert.Contains(t, actual, "`a`=?")
	assert.Contains(t, actual, "AND `b`=?")
}

func TestParserSelectUnchanged(t *testing.T) {
	testParser(t,
		`
<mapper namespace="Test">
    <select id="testSelect">
        SELECT name FROM fruits WHERE id = #{id}
    </select>
</mapper>`,
		"SELECT `name` FROM `fruits` WHERE `id`=?;",
	)
}

func xmlTestDataDir(t *testing.T) string {
	t.Helper()
	_, file, _, ok := runtime.Caller(0)
	if !assert.True(t, ok) {
		t.FailNow()
	}
	return filepath.Join(filepath.Dir(file), "testdata", "xml")
}

func readXMLTestFile(t *testing.T, name string) string {
	t.Helper()
	content, err := os.ReadFile(filepath.Join(xmlTestDataDir(t), name))
	if !assert.NoError(t, err) {
		t.FailNow()
	}
	return string(content)
}

func findStmtByStartLine(stmts []ast.StmtInfo, startLine uint64) (ast.StmtInfo, bool) {
	for _, stmt := range stmts {
		if stmt.StartLine == startLine {
			return stmt, true
		}
	}
	return ast.StmtInfo{}, false
}

func TestParseXMLsRegressionSetWhereMapper(t *testing.T) {
	content := readXMLTestFile(t, "regression_set_where_mapper.xml")
	stmts, err := ParseXMLs([]XmlFile{{FilePath: "regression_set_where_mapper.xml", Content: content}})
	if !assert.NoError(t, err) {
		return
	}
	assert.Len(t, stmts, 23)

	updateToReverse, ok := findStmtByStartLine(stmts, 1294)
	if !assert.True(t, ok) {
		return
	}
	assert.NotContains(t, updateToReverse.SQL, "nullWHERE")
	assert.Contains(t, updateToReverse.SQL, "NULL")
	assert.Contains(t, updateToReverse.SQL, "WHERE")

	for _, stmt := range stmts {
		assert.NotContains(t, stmt.SQL, "nullWHERE")
	}
}

func TestParseXMLsRegressionBasicMapper(t *testing.T) {
	content := readXMLTestFile(t, "regression_basic_mapper.xml")
	stmts, err := ParseXMLs([]XmlFile{{FilePath: "regression_basic_mapper.xml", Content: content}})
	if !assert.NoError(t, err) {
		return
	}
	assert.Len(t, stmts, 10)

	for _, stmt := range stmts {
		assert.NotEmpty(t, strings.TrimSpace(stmt.SQL))
		assert.NotContains(t, stmt.SQL, "nullWHERE")
	}

	testParameters, ok := findStmtByStartLine(stmts, 16)
	if assert.True(t, ok) {
		assert.Contains(t, testParameters.SQL, "`category`=?")
		assert.Contains(t, testParameters.SQL, "`price`>?")
	}

	testChoose, ok := findStmtByStartLine(stmts, 96)
	if assert.True(t, ok) {
		assert.Contains(t, testChoose.SQL, "`category`=\"apple\"")
	}
}
