package parser

import "testing"

func TestIbatis(t *testing.T) {
	testParser(t, `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE sqlMap PUBLIC "-//ibatis.apache.org//DTD SQL Map 2.0//EN" "http://ibatis.apache.org/dtd/sql-map-2.dtd">

<sqlMap namespace="Employee">

   <select id="findByID" resultClass="Employee">
      SELECT * FROM EMPLOYEE
		
      <dynamic prepend="WHERE ">
         <isNull property="id" prepend="AND ">
            id IS NULL
         </isNull>
         <isNotNull property="id" prepend="OR ">
            id = #id#
         </isNotNull>
      </dynamic>
		
   </select>
</sqlMap>`,
		"SELECT * FROM `EMPLOYEE` WHERE `id` IS NULL OR `id`=?;")

	testParser(t, `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE sqlMap PUBLIC "-//ibatis.apache.org//DTD SQL Map 2.0//EN" "http://ibatis.apache.org/dtd/sql-map-2.dtd">

<sqlMap namespace="Employee">

<select id="findByID" resultClass="Employee">
SELECT * FROM EMPLOYEE

<dynamic prepend="WHERE ">
<iterate prepend="AND" property="UserNameList"
open="(" close=")" conjunction="OR">
username=#UserNameList[]#
</iterate>
<isNull property="id">
id IS NULL
</isNull>

<isNotNull property="id">
id = #id#
</isNotNull>
</dynamic>

</select>
</sqlMap>`,
		"SELECT * FROM `EMPLOYEE` WHERE (`username`=? OR `username`=?) AND `id` IS NULL AND `id`=?;")
}

func TestParseIBatisInclude(t *testing.T) {
	testParserQuery(t, false, `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE sqlMap PUBLIC "-//ibatis.apache.org//DTD SQL Map 2.0//EN" "http://ibatis.apache.org/dtd/sql-map-2.dtd">

<sqlMap namespace="Employee">

<sql id="selectItem_fragment">
FROM items
WHERE parentid = 6
</sql>

<select id="selectItemCount" resultClass="int">
SELECT COUNT(*) AS total
<include refid="selectItem_fragment"/>
</select>
<select id="selectItems" resultClass="Item">
SELECT id, name
<include refid="selectItem_fragment"/>
</select>
</sqlMap>`,
		[]string{
			"SELECT COUNT(1) AS `total` FROM `items` WHERE `parentid`=6",
			"SELECT `id`,`name` FROM `items` WHERE `parentid`=6",
		})

	testParserQuery(t, false, `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE sqlMap PUBLIC "-//ibatis.apache.org//DTD SQL Map 2.0//EN" "http://ibatis.apache.org/dtd/sql-map-2.dtd">

<sqlMap namespace="Employee">

<sql id="selectItem_fragment">
FROM items
WHERE parentid = #value#
</sql>
<select id="selectItemCount" parameterClass="int" resultClass="int">
SELECT COUNT(*) AS total
<include refid="selectItem_fragment"/>
</select>
<select id="selectItems" parameterClass="int" resultClass="Item">
SELECT id, name
<include refid="selectItem_fragment"/>
</select>

</sqlMap>`,
		[]string{
			"SELECT COUNT(1) AS `total` FROM `items` WHERE `parentid`=?",
			"SELECT `id`,`name` FROM `items` WHERE `parentid`=?",
		})
}

func TestParseIBatisAll(t *testing.T) {
	testParserQuery(t, true, `
<!DOCTYPE sqlMap PUBLIC "-//ibatis.apache.org//DTD SQL Map 2.0//EN" "http://ibatis.apache.org/dtd/sql-map-2.dtd">

<sqlMap namespace="Employee">
    <sql id="selectItem_fragment">
        FROM Employee
        WHERE id = 6
    </sql>

    <sql id="selectItem_fragment2">
        FROM Employee
        WHERE id = #value#
    </sql>

    <select id="selectItemCount1" resultClass="int">
        SELECT COUNT(*) AS total
        <include refid="selectItem_fragment"/>
    </select>

    <select id="selectItems1" resultClass="Item">
        SELECT id, name
        <include refid="selectItem_fragment"/>
    </select>

    <select id="selectItemCount2" parameterClass="int" resultClass="int">
        SELECT COUNT(*) AS total
        <include refid="selectItem_fragment2"/>
    </select>

    <select id="selectItems2" parameterClass="int" resultClass="Item">
        SELECT id, name
        <include refid="selectItem_fragment2"/>
    </select>

    <select id="findByID1" resultClass="Employee">
        SELECT * FROM EMPLOYEE
        <dynamic prepend="WHERE ">
            <isNull property="id" prepend="AND ">
                id IS NULL
            </isNull>
            <isNotNull property="id" prepend="OR ">
                id = #id#
            </isNotNull>
        </dynamic>
    </select>

    <select id="findByID2" resultClass="Employee">
        SELECT * FROM EMPLOYEE
        <dynamic prepend="WHERE ">
            <iterate prepend="AND" property="UserNameList" open="(" close=")" conjunction="OR">
                username=#UserNameList[]#
            </iterate>
            <isNull property="id">
                id IS NULL
            </isNull>
            <isNotNull property="id">
                id = #id#
            </isNotNull>
        </dynamic>
    </select>

    <select id="dynamicGetAccountList" resultMap="account-result" >
        select * from EMPLOYEE
        <dynamic prepend="WHERE">
            <isNotNull prepend="AND" property="firstName" open="(" close=")">
                ACC_FIRST_NAME = #firstName#
                <isNotNull prepend="OR" property="lastName">
                    ACC_LAST_NAME = #lastName#
                </isNotNull>
            </isNotNull>
            <isNotNull prepend="AND" property="emailAddress">
                ACC_EMAIL like #emailAddress#
            </isNotNull>
            <isGreaterThan prepend="AND" property="id" compareValue="0">
                ACC_ID = #id#
            </isGreaterThan>
        </dynamic>
        order by ACC_LAST_NAME
    </select>

    <select id="getProduct" resultMap="get-product-result">
        select * from EMPLOYEE order by $preferredOrder$
    </select>
</sqlMap>
`, []string{
		"SELECT COUNT(1) AS `total` FROM `Employee` WHERE `id`=6",
		"SELECT `id`,`name` FROM `Employee` WHERE `id`=6",
		"SELECT COUNT(1) AS `total` FROM `Employee` WHERE `id`=?",
		"SELECT `id`,`name` FROM `Employee` WHERE `id`=?",
		"SELECT * FROM `EMPLOYEE` WHERE `id` IS NULL OR `id`=?",
		"SELECT * FROM `EMPLOYEE` WHERE (`username`=? OR `username`=?) AND `id` IS NULL AND `id`=?",
		"SELECT * FROM `EMPLOYEE` WHERE `ACC_FIRST_NAME`=? OR `ACC_LAST_NAME`=? AND `ACC_EMAIL` LIKE ? AND `ACC_ID`=? ORDER BY `ACC_LAST_NAME`",
		"SELECT * FROM `EMPLOYEE` ORDER BY ?",
	})
}