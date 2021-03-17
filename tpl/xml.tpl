<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE mapper PUBLIC "-//mybatis.org//DTD Mapper 3.0//EN" "http://mybatis.org/dtd/mybatis-3-mapper.dtd">
<!--
 {{if .Config.XML.Comment}}{{.XML.Mapper.Entity.Table.Comment}} Mapper XML{{end}}
 @author {{.Config.Global.Author}}
 {{if .Config.Global.Date}}@since {{.Extra.Date}}{{end}}
 {{if .Config.Global.Copyright}}@created by {{.Config.Global.CopyrightContent}}{{end}}
 {{if .Config.Global.Website}}@repo {{.Config.Global.WebsiteContent}}{{end}}
-->
<mapper namespace="{{.XML.Mapper.PKGName}}">
    <resultMap id="BaseResultMap" type="{{.XML.Mapper.Entity.PKGName}}">{{range $i, $e := .XML.Mapper.Entity.Ids}}
        <id column="{{$e.Column.Name}}" jdbcType="{{$e.Column.UpperType}}" property="{{$e.Name}}"/>{{end}}{{range $i, $e := .XML.Mapper.Entity.Fields}}
        <result column="{{$e.Column.Name}}" jdbcType="{{$e.Column.UpperType}}" property="{{$e.Name}}"/>{{end}}
    </resultMap>
</mapper>