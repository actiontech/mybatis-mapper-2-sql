package main

import (
	"fmt"
	"testing"
)

var a = `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE mapper PUBLIC "-//mybatis.org//DTD Mapper 3.0//EN" "http://mybatis.org/dtd/mybatis-3-mapper.dtd">
<mapper namespace="Test">
    <sql id="someinclude">
        FROM
        <include refid="${include_target}"/>
        <include refid="somewhere"/>
    </sql>
    <update id="testSet">
        UPDATE
        fruits
        <set>
            <if test="category != null and category !=''">
                category = #{category},
            </if>
            <if test="price != null and price !=''">
                price = ${price},
            </if>
        </set>
        WHERE
        name = #{name}
    </update>
</mapper>
`

func TestParser(t *testing.T) {
	s, err := Parse(a)
	if err != nil {
		t.Errorf("parse error: %v", err)
		return
	}
	fmt.Println(s)
}
