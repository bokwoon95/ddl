-- Get tables

SELECT tbl_name AS table_name
FROM sqlite_schema
WHERE type = 'table';

-- Get columns

WITH unique_columns AS (
    SELECT
        ii.cid
        ,ii.name
    FROM
        pragma_index_list('customer') AS il, pragma_index_info(il.name) AS ii
    WHERE
        il."unique" = TRUE
        AND il.origin = 'u'
        AND (
            SELECT COUNT(*)
            FROM pragma_index_list('customer') AS il2, pragma_index_info(il.name) AS ii2
            WHERE il2.name = il.name
            GROUP BY il2.name
        ) = 1
)
SELECT
    'customer' AS table_name -- TableName
    ,ti.name AS column_name -- ColumnName
    ,ti."type" AS column_type -- Type
    ,ti."notnull" AS not_null -- NotNull
    ,ti.pk AS is_primary_key -- IsPrimaryKey
    ,CASE WHEN unique_columns.name IS NULL THEN FALSE ELSE TRUE END AS is_unique -- IsUnique
    ,ti."type" = 'INTEGER' AND ti.pk AS is_autoincrement -- IsAutoincrement
    ,ti.dflt_value AS column_default -- ColumnDefault
    ,fkl."table" AS references_table -- ReferencesTable
    ,fkl."to" AS references_column -- ReferencesColumn
    ,fkl.on_update AS references_on_update -- ReferencesOnUpdate
    ,fkl.on_delete AS references_on_delete -- ReferencesOnDelete
FROM
    pragma_table_info('customer') AS ti
    LEFT JOIN  unique_columns ON unique_columns.cid = ti.cid
    LEFT JOIN pragma_foreign_key_list('customer') AS fkl ON fkl."from" = ti.name
;

WITH unique_columns AS (
    SELECT
        ii.cid
        ,ii.name
    FROM
        pragma_index_list('customer') AS il, pragma_index_info(il.name) AS ii
    WHERE
        il."unique" = TRUE
        AND il.origin = 'u'
        AND (
            SELECT COUNT(*)
            FROM pragma_index_list('customer') AS il2, pragma_index_info(il.name) AS ii2
            WHERE il2.name = il.name
            GROUP BY il2.name
        ) = 1
)
SELECT
    'customer' AS table_name -- TableName
    ,ti.name AS column_name -- ColumnName
    ,ti."type" AS column_type -- Type
    ,ti."notnull" AS not_null -- NotNull
    ,ti.pk AS is_primary_key -- IsPrimaryKey
    ,CASE WHEN unique_columns.name IS NULL THEN FALSE ELSE TRUE END AS is_unique -- IsUnique
    ,ti."type" = 'INTEGER' AND ti.pk AS is_autoincrement -- IsAutoincrement
    ,ti.dflt_value AS column_default -- ColumnDefault
    ,fkl."table" AS references_table -- ReferencesTable
    ,fkl."to" AS references_column -- ReferencesColumn
    ,fkl.on_update AS references_on_update -- ReferencesOnUpdate
    ,fkl.on_delete AS references_on_delete -- ReferencesOnDelete
FROM
    pragma_table_info('dummy_table') AS ti
    LEFT JOIN  unique_columns ON unique_columns.cid = ti.cid
    LEFT JOIN pragma_foreign_key_list('dummy_table') AS fkl ON fkl."from" = ti.name
;

-- Get constraints


-- Get indices

WITH indexed_columns AS (
    SELECT
        'customer' AS table_name
        ,il.name AS index_name
        ,il."unique" AS is_unique
        ,il.partial AS is_partial
        ,CASE ii.cid WHEN -1 THEN '' WHEN -2 THEN '' ELSE ii.name END AS column_name
    FROM
        pragma_index_list('customer') AS il
        CROSS JOIN pragma_index_info(il.name) AS ii
    WHERE
        il.origin <> 'u'
    ORDER BY
        il.name
        ,ii.seqno
)
SELECT
    table_name -- TableName
    ,index_name -- IndexName
    ,is_unique -- IsUnique
    ,is_partial -- IsPartial
    ,json_group_array(column_name) AS columns -- Columns
FROM
    indexed_columns
GROUP BY
    table_name
    ,index_name
    ,is_unique
    ,is_partial
;
