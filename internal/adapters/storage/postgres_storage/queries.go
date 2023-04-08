package postgres_storage

var (
	createLayer      = `call public.create_layer(:layer_name);`
	selectLayerNames = `select substring(schema_name, 1, position('_layer' in schema_name) - 1) 
							 from information_schema.schemata where schema_name like '%_layer';`

	insertUser         = `select * from public.insert_user(:username, :password, :role);`
	selectUserPassword = `select password from public.users where name = $1 limit 1;`
	selectRole         = `select users.role from public.users where users.name = (select session_user);`

	selectAllModels = `select * from public.select_all_models(:layer_name);`
	insertModels    = `select * from public.insert_models(:layer_name, :models_array);`

	selectAllModelElements = `select * from public.select_all_model_elements(:layer_name);`
	insertModelElements    = `select * from public.insert_model_elements(:layer_name, :elements_array);`

	selectAllProperties        = `select * from public.properties;`
	selectPropertiesByUnit     = `select * from public.select_properties_by_unit(:layer_name, :lang, :unit_text);`
	selectPropertiesIdByUnitId = `select * from public.select_properties_id_by_unit_id(:layer_name, :lang, :unit_id);`
	insertProperties           = `select * from public.insert_properties(:properties_array);`

	selectContextsById     = `select id, text from public.contexts where id = any($1);`
	selectContextsIdByUnit = `select * from public.select_contexts_id_by_unit(:layer_name, :lang, :unit_id);`

	selectAllLinkedUnits                     = `select * from public.select_all_linked_units(:layer_name);`
	selectLinkedUnitsByModelsId              = `select * from public.select_linked_units_by_models_id(:layer_name, :models_id_array);`
	selectLinkedUnitsByPropertiesId          = `select * from public.select_linked_units_by_properties_id(:layer_name, :properties_id_array);`
	selectUnlinkedUnitsByLang                = `select * from public.select_unlinked_units_by_lang(:layer_name, :lang);`
	selectUnlinkedUnitsByLangAndModelsId     = `select * from public.select_unlinked_units_by_lang_and_models_id(:layer_name, :lang, :models_id_array);`
	selectUnlinkedUnitsByLangAndPropertiesId = `select * from public.select_unlinked_units_by_lang_and_properties_id(:layer_name, :lang, :properties_id_array);`
)
