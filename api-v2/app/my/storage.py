
import os
from uuid import uuid4
from eve.io.media import MediaStorage
from flask import current_app as app
from google.cloud import storage


class GoogleMediaStorage(MediaStorage):
    def put(self, content, filename=None, content_type=None, resource=None):
        if resource in ('modules', 'users'):
            ext = filename.rsplit('.', 1)[1].lower()
            name = uuid4().hex+'.'+ext
            # name = secure_filename(content.filename)
            location = os.path.join(app.config['UPLOAD_FOLDER'], name)
            content.save(location)
            client = storage.Client.from_service_account_json(
                app.config['MONGO_JSON'])
            bucket = client.bucket(app.config['MEDIA_ENDPOINT'])
            blob = bucket.blob(name)

            blob.upload_from_filename(location, content_type=content_type)

            return name

    def get(self, id_or_filename, resource=None):
        if resource in ('modules', 'users'):
            return id_or_filename

    def delete(self, id_or_filename, resource=None):
        pass
