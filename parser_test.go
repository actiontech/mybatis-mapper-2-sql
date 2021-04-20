package main

import (
	"testing"
)

func TestParserInclude(t *testing.T) {
	testParser(t,
		`
<mapper namespace="Test">
	<sql id="sometable">
  		${prefix}Table
	</sql>
	<sql id="someinclude">
  		from
    	<include refid="${include_target}"/>
	</sql>
	<select id="select" resultType="map">
  		select
    	field1, field2, field3
  		<include refid="someinclude">
    		<property name="prefix" value="Some"/>
    		<property name="include_target" value="sometable"/>
  		</include>
	</select>
</mapper>`,
`
		select
		field1, field2, field3
		
		from
		
		SomeTable
`)
}

func TestParserIf(t *testing.T) {
	testParser(t,
		`
<mapper namespace="Test">
	   <select id="testIf">
        SELECT
        name,
        category,
        price
        FROM
        fruits
        WHERE
        1=1
        <if test="category != null and category !=''">
            AND category = #{category}
        </if>
        <if test="price != null and price !=''">
            AND price = ${price}
            <if test="price >= 400">
                AND name = 'Fuji'
            </if>
        </if>
    </select>
</mapper>`,
		`
		select
		field1, field2, field3
		
		from
		
		SomeTable
`)
}


func testParser(t *testing.T, xmlData, expect string) {
	actual, err := ParseXML(xmlData)
	if err != nil {
		t.Errorf("parse error: %v", err)
		return
	}
	if actual != expect {
		t.Errorf("\nexpect: [%s]\nactual: [%s]", expect, actual)
	}
}