from pprint import pprint
import os
from uuid import uuid4

from eve.io.media import MediaStorage
from flask import current_app as app, Response, abort
from eve.auth import TokenAuth
from tables import Users
from minio import Minio
from minio.error import ResponseError
from eve_sqlalchemy.validation import ValidatorSQL
from werkzeug.datastructures import FileStorage
from werkzeug.utils import secure_filename

minioClient = Minio('%s:9000' % os.environ['HOST_IP'],
                    access_key=os.environ['MINIO_ACCESS_KEY'],
                    secret_key=os.environ['MINIO_SECRET_KEY'],
                    secure=False)


class MyAuth(TokenAuth):

    def authenticate(self):
        resp = Response(None, 401)
        abort(401, description='Please provide proper credentials',
              response=resp)

    def check_auth(self, token, allowed_roles, resource, method):
        # pprint(allowed_roles)
        # pprint(resource)
        # pprint(method)
        user = Users.auth(token)

        if len(user):
            self.set_request_auth_value(user['id'])
        # pprint(app.auth.get_request_auth_value())
        # current_app.auth.get_request_auth_value()

        return len(user)


class MyMediaStorage(MediaStorage):
    def put(self, content, filename=None, content_type=None, resource=None):
        if resource in ('modules'):
            # pprint(content)
            # pprint(filename)
            # pprint(content_type)
            ext = filename.rsplit('.', 1)[1].lower()
            name = uuid4().hex+'.'+ext
            # name = secure_filename(content.filename)
            location = os.path.join(app.config['UPLOAD_FOLDER'], name)
            content.save(location)
            minioClient.fput_object(
                'module-images', name, location, content_type)
            url = minioClient.presigned_get_object('module-images', name)
            return name
            # return url.split('?')[0].split('/')[-1]
        return ''

    def get(self, id_or_filename, resource=None):
        if resource in ('modules'):
            # pprint(id_or_filename)
            # pprint(resource)
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

    # def _normalize_coerce_upload(self, value):
    #     # db_images = app.data.driver.db['images']
    #     name = app.media.put(value, filename=value.filename,
    #                          content_type=value.mimetype, resource='modules')

    #     # date_utc = datetime.utcnow().replace(microsecond=0)
    #     # documents = [{
    #     #     'owner': ObjectId(app.auth.get_request_auth_value()),
    #     #     'image': name,
    #     #     app.config['LAST_UPDATED']: date_utc,
    #     #     app.config['DATE_CREATED']: date_utc,
    #     # }]
    #     # resolve_document_etag(documents, 'images')
    #     # image_id = db_images.insert_one(documents[0]).inserted_id

    #     return name
