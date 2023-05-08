#!/usr/bin/python3
# -*- coding: utf-8 -*-
import requests

def print_request(request):
    req = "{method} {path_url} HTTP/1.1\r\n{headers}\r\n{body}".format(
        method = request.method,
        path_url = request.path_url,
        headers = ''.join('{0}: {1}\r\n'.format(k, v) for k, v in request.headers.items()),
        body = request.body or "",
    )
    return "{req_size}\n{req}\r\n".format(req_size = len(req), req = req)

#POST multipart form data
def post_multipart(host, port, namespace, files, headers, payload, json):
    req = requests.Request(
        'PUT',
        'http://{host}:{port}{namespace}'.format(
            host = host,
            port = port,
            namespace = namespace,
        ),
        headers = headers,
        data = payload,
        files = files,
        json = json
    )
    prepared = req.prepare()
    return print_request(prepared)

if __name__ == "__main__":
    #usage sample below
    #target's hostname and port
    #this will be resolved to IP for TCP connection
    host = 'localhost'
    port = '8080'
    namespace = '/api/v1/admin/units/all?layer=newlayer&token=10063865700249539947'
    #below you should specify or able to operate with
    #virtual server name on your target
    headers = {
        'Host': '127.0.0.1'
    }
    payload = {
    }
    files = {
        # name, path_to_file, content-type, additional headers
    }
    json = {
    }

    print(post_multipart(host, port, namespace, files, headers, payload, json))
