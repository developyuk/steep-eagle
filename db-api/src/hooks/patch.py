import os
from pprint import pprint

from flask import current_app as app, jsonify
from werkzeug import secure_filename
from minio import Minio
from minio.error import ResponseError

minioClient = Minio('%s:9000' % os.environ['HOST_IP'],
                    access_key=os.environ['MINIO_ACCESS_KEY'],
                    secret_key=os.environ['MINIO_SECRET_KEY'],
                    secure=False)


def module_image(name, location,type):
    url = ''
    try:
        minioClient.fput_object('modules-img', name, location,type)
        return minioClient.presigned_put_object('modules-img', name)
    except ResponseError as err:
        return jsonify({"_status": "ERR", "_error": {'message': str(e), "code": 400}}), 400
    return url


def before_patch_item(resource_name, updates, original):
    data = updates
    image = data.get('fimage')
    if resource_name == 'modules' and image:
        filename = secure_filename(image.filename)
        location = os.path.join(app.config['UPLOAD_FOLDER'], filename)
        image.save(location)
        url = module_image(filename, location,image.content_type)
        url = url.split('?')[0]
        updates.update({'image': url})
