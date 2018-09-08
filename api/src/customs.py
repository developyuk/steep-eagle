
import os
from uuid import uuid4
import json
import datetime

from eve.io.media import MediaStorage
from flask import current_app as app, Response, abort
from eve.auth import TokenAuth
from minio import Minio

from eve_sqlalchemy.validation import ValidatorSQL
from werkzeug.datastructures import FileStorage
# from werkzeug.utils import secure_filename
import jwt

minioClient = Minio('%s:9000' % os.environ['HOST_IP'],
                    access_key=os.environ['MINIO_ACCESS_KEY'],
                    secret_key=os.environ['MINIO_SECRET_KEY'],
                    secure=False)


def auth(token):
    try:
        data = jwt.decode(token, app.config['JWT_SECRET'])
    except Exception as e:
        raise Exception(str(e))

    return data


class MyAuth(TokenAuth):

    def authenticate(self):
        resp = Response(None, 401)
        abort(401, description='Please provide proper credentials', response=resp)

    def check_auth(self, token, allowed_roles, resource, method):
        # pprint(allowed_roles)
        # pprint(resource)
        # pprint(method)
        user = auth(token)

        if len(user):
            self.set_request_auth_value(user['id'])
        # pprint(app.auth.get_request_auth_value())

        return len(user)


class MyMediaStorage(MediaStorage):
    def put(self, content, filename=None, content_type=None, resource=None):
        if resource in ('modules', 'users'):
            ext = filename.rsplit('.', 1)[1].lower()
            name = uuid4().hex+'.'+ext
            # name = secure_filename(content.filename)
            location = os.path.join(app.config['UPLOAD_FOLDER'], name)
            content.save(location)
            minioClient.fput_object(
                app.config['MEDIA_ENDPOINT'], name, location, content_type)
            url = minioClient.presigned_get_object(
                app.config['MEDIA_ENDPOINT'], name)
            return name
        return ''

    def get(self, id_or_filename, resource=None):
        if resource in ('modules', 'users'):
            return id_or_filename

    def delete(self, id_or_filename, resource=None):
        pass


class MyValidator(ValidatorSQL):
    def _validate_type_media(self, field, value):
        """ Enables validation for `media` data type.

        :param field: field name.
        :param value: field value.

        .. versionadded:: 0.3
        """
        if not isinstance(value, FileStorage):
            self._error(field, "file was expected, got '%s' instead." % value)
