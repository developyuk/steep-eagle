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


def module_image(image):
    try:
        filename = secure_filename(image.filename)
        location = os.path.join(app.config['UPLOAD_FOLDER'], filename)
        image.save(location)
        minioClient.fput_object('modules-img', filename,
                                location, image.content_type)
        url = minioClient.presigned_put_object('modules-img', filename)
        return url.split('?')[0]
    except Exception as e:
        return jsonify({"_status": "ERR", "_error": {'message': str(e), "code": 400}}), 400


def before_patch_item(resource_name, updates, original):
    image = updates.get('fimage')
    if resource_name == 'modules' and image:
        url = module_image(image)
        updates.update({'image': url})


def before_post_item(resource_name, items):
    pprint(resource_name)
    for i, item in enumerate(items):
        image = item.get('fimage')
        if resource_name == 'modules' and image:
            url = module_image(image)
            items[i].update({'image': url})
