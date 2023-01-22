namespace go toutiao.middleware.hzClient

struct FormReq {
    1: string FormValue (api.form="form1");
    2: string FileValue (api.file_name="file1");
}

struct QueryReq {
    1: string QueryValue (api.query="query1");
}

struct PathReq {
    1: string PathValue (api.path="path1");
}

struct BodyReq {
    1: string BodyValue (api.body="");
    2: string QueryValue (api.query="query2");
}

struct Resp {
    1: string Resp;
}

service Hertz {
    Resp FormMethod(1: FormReq request) (api.post="/form", api.handler_path="post");
    Resp QueryMethod(1: QueryReq request) (api.get="/query", api.handler_path="get");
    Resp PathMethod(1: PathReq request) (api.post="/path:path1", api.handler_path="post");
    Resp BodyMethod(1: BodyReq request) (api.post="/body", api.handler_path="post");
}(
    api.base_domain="http://127.0.0.1:8888";
)