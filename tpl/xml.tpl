<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE mapper PUBLIC "-//mybatis.org//DTD Mapper 3.0//EN" "http://mybatis.org/dtd/mybatis-3-mapper.dtd">
<mapper namespace="{{.XML.Mapper.PKGName}}">
    {{range $i, $e := .Config.XML.MapperPrefixes}}
    {{$e}}
    {{end}}

    <resultMap id="BaseResultMap" type="{{.XML.Mapper.Entity.PKGName}}">
        {{.XML.ResultMap.HaveId}}<id column="{{.XML.ResultMap.Id.Field.Column.Name}}" jdbcType="{{.XML.ResultMap.Id.Field.Column.UpperType}}" property="{{.XML.ResultMap.Id.Field.Name}}"/>{{end}}
        {{range $i, $e := .XML.ResultMap.Items}}
        <result column="{{$e.Field.Column.Name}}" jdbcType="{{$e.Field.Column.UpperType}}" property="{{$e.Field.Name}}"/>
        {{end}}
    </resultMap>

    {{range $i, $e := .Config.XML.MapperSuffixes}}
    {{$e}}
    {{end}}
</mapper>