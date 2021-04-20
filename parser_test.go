package main

import (
	"testing"
)

func TestParser(t *testing.T) {
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