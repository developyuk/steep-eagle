from pprint import pprint
import os
from eve.io.media import MediaStorage

from flask import current_app as app, Response, abort
from eve.auth import TokenAuth
from tables import Users
from minio import Minio
from minio.error import ResponseError
from eve_sqlalchemy.validation import ValidatorSQL

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

            name = secure_filename(content.filename)
            location = os.path.join(app.config['UPLOAD_FOLDER'], name)
            content.save(location)
            minioClient.fput_object('modules-img', name, location, content_type)
            url = minioClient.presigned_get_object('modules-img', name)
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
    # def _validate_isodd(self, isodd, field, value):
    #     if isodd and not bool(value & 1):
    #         self._error(field, "Value must be an odd number")
    # def _validate_type_ghost(self, value):
    # """ Enables validation for `objectid` schema attribute.

    # :param value: field value.
    # """
    # if isinstance(value, ObjectId):
    #     return True
    def _normalize_coerce_upload(self, value):
        db_images = app.data.driver.db['images']
        name = app.media.put(value, filename=value.filename,
                             content_type=value.mimetype, resource='images')

        # date_utc = datetime.utcnow().replace(microsecond=0)
        # documents = [{
        #     'owner': ObjectId(app.auth.get_request_auth_value()),
        #     'image': name,
        #     app.config['LAST_UPDATED']: date_utc,
        #     app.config['DATE_CREATED']: date_utc,
        # }]
        # resolve_document_etag(documents, 'images')
        # image_id = db_images.insert_one(documents[0]).inserted_id

        return name
