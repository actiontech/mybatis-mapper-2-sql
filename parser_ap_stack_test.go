package parser

import "testing"

func TestDynamicSQL(t *testing.T) {
	testParserQuery(t, true, `
<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<sqls id="Bc" longname="批量SQL" package="cn.tt"
      xsi:noNamespaceSchemaLocation="ltts-model.xsd" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
    <dynamicSelect cache="none" method="selectAll" type="sql" id="test1" longname="test">
        <dynamicSql type="mysql">
            <str type="Str"><![CDATA[
    	select tranno from rps_errr_detl where ckdate = #ckdate# and hadres in (3,7)
    	]]></str>
        </dynamicSql>
    </dynamicSelect>

    <dynamicUpdate method="update" id="test" longname="test">
        <dynamicSql type="mysql">
            <str type="Str"><![CDATA[
				UPDATE rps_mesg_push p SET p.msgsta=4,p.gmt_modified=now(6) WHERE p.msgsta IN (1,3) AND p.temcod=${temcod}
			]]></str>
        </dynamicSql>
    </dynamicUpdate>

    <dynamicDelete method="delete" id="delTableDaByid" longname="test">
        <dynamicSql type="mysql">
            <str type="Str"><![CDATA[delete from ${tablena} ]]></str>
            <where type="Where">
                <str type="Str"><![CDATA[id between #minid# and #maxid# limit ${limitcout} ]]></str>
            </where>
        </dynamicSql>
    </dynamicDelete>

    <dynamicSelect cache="none" method="selectOne" type="sql" id="selWarnInfoDa" longname="test">
        <dynamicSql type="mysql">
            <if test="schema!=null &amp;&amp; schema!=&quot;&quot;" type="If">
                <str type="Str"><![CDATA[${schema}]]></str>
            </if>
    		<if test="schema!=null &amp;&amp; schema!=&quot;&quot;" type="If">
                <str type="Str"><![CDATA[${schema}]]></str>
            </if>
            <str type="Str"><![CDATA[select count(*) from rps_warn_info ]]></str>
            <where type="Where">
                <and type="And">
                    <if test="warnty!=&quot;&quot;&amp;&amp;warnty!=null" type="If">
                        <str type="Str"><![CDATA[warnty in (${warnty})]]></str>
                    </if>
                    <str type="Str"><![CDATA[gmt_modified between #starttime# and #endtime#]]></str>
                </and>
            </where>
        </dynamicSql>
    </dynamicSelect>
</sqls>

`, []string{
		"SELECT `tranno` FROM `rps_errr_detl` WHERE `ckdate`=? AND `hadres` IN (3,7)",
		"UPDATE `rps_mesg_push` AS `p` SET `p`.`msgsta`=4, `p`.`gmt_modified`=now(6) WHERE `p`.`msgsta` IN (1,3) AND `p`.`temcod`=?",
		" delete from ?  WHERE id between ? and ? limit ?",
		"SELECT count(1) FROM `rps_warn_info` WHERE (`warnty` IN (?)) AND (`gmt_modified` BETWEEN ? AND ?)",
	})
}
