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
