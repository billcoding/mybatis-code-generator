<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE mapper PUBLIC "-//mybatis.org//DTD Mapper 3.0//EN" "http://mybatis.org/dtd/mybatis-3-mapper.dtd">
<mapper namespace="{{.XML.Mapper.PKGName}}">
    {{range $i, $e := .Config.Entity.MapperPrefixes}}
    {{$e}}
    {{end}}

    <resultMap id="BaseResultMap" type="{{.XML.Entity.PKGName}}">
        {{.XML.ResultMap.HaveId}}<id column="{{.XML.ResultMap.Id.Column}}" jdbcType="{{.XML.ResultMap.Id.MysqlType}}" property="{{.XML.ResultMap.Id.Name}}"/>{{end}}
        {{range $i, $e := ..XML.ResultMap.Items}}
        <result column="{{$e.Column.Name}}" jdbcType="{{$e.Column.MysqlType}}" property="{{$e.Field.Name}}"/>
        {{end}}
    </resultMap>

    {{range $i, $e := .Config.XML.MapperPrefixes}}
    {{$e}}
    {{end}}
</mapper>