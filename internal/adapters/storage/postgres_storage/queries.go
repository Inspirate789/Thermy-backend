package postgres_storage

var (
	createLayerQuery      = `call public.create_layer(:layer_name);`
	selectLayerNamesQuery = `select substring(schema_name, 1, position('_layer' in schema_name) - 1) 
							 from information_schema.schemata where schema_name like '%_layer';`

	insertUserQuery         = `select * from public.insert_user(:username, :password, :role);`
	selectRoleQuery         = `select users.role from public.users where users.name = (select session_user);`
	selectUserPasswordQuery = `select password from public.users where name = $1 limit 1;`

	selectAllModelsQuery = `select * from public.select_all_models(:layer_name);`

	selectAllModelElementsQuery = `select * from public.select_all_model_elements(:layer_name);`

	selectAllPropertiesQuery        = `select * from public.properties;`
	selectPropertiesByUnitQuery     = `select * from public.select_properties_by_unit(:layer_name, :lang, :unit_text);`
	selectPropertiesIdByUnitIdQuery = `select * from public.select_properties_id_by_unit_id(:layer_name, :lang, :unit_id);`

	selectContextsByIdQuery     = `select id, text from public.contexts where id = any($1);`
	selectContextsIdByUnitQuery = `select * from public.select_contexts_id_by_unit(:layer_name, :lang, :unit_id);`

	selectAllLinkedUnitsQuery      = `select * from public.select_all_linked_units(:layer_name);`
	selectUnlinkedUnitsByLangQuery = `select * from public.select_unlinked_units_by_lang(:layer_name, :lang);`
)
