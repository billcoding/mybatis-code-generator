<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE mapper PUBLIC "-//mybatis.org//DTD Mapper 3.0//EN" "http://mybatis.org/dtd/mybatis-3-mapper.dtd">
<mapper namespace="{{.XML.Mapper.PKGName}}">
    {{range $i, $e := .Config.XML.MapperPrefixes}}{{$e}}
    {{end}}
    <resultMap id="BaseResultMap" type="{{.XML.Mapper.Entity.PKGName}}">
        {{if .XML.ResultMap.HaveId}}<id column="{{.XML.ResultMap.Id.Column.Name}}" jdbcType="{{.XML.ResultMap.Id.Column.UpperType}}" property="{{.XML.ResultMap.Id.Name}}"/>{{end}}
        {{range $i, $e := .XML.ResultMap.Items}}<result column="{{$e.Column.Name}}" jdbcType="{{$e.Column.UpperType}}" property="{{$e.Name}}"/>
        {{end}}
    </resultMap>
    {{range $i, $e := .Config.XML.MapperSuffixes}}{{$e}}
    {{end}}
</mapper>