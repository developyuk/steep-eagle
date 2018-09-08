from pytz import timezone, utc
from datetime import timedelta, datetime

from flask import current_app as app, abort
from eve.utils import parse_request
from eve.methods.common import ratelimit, epoch, pre_event, resolve_embedded_fields, \
    build_response_document, resource_link, document_link, last_updated

dow = dict(
    zip('monday tuesday wednesday thursday friday saturday sunday'.split(), range(7)))

utc_now = utc.localize(datetime.utcnow())
wib_now = utc_now.astimezone(timezone("Asia/Jakarta"))


def get_internal(resource, **lookup):
    response = {}
    documents = []
    req = parse_request(resource)
    embedded_fields = resolve_embedded_fields(resource, req)
    # req.show_deleted = True
    cursor = app.data.find(resource, req, lookup)
    for document in cursor:
        build_response_document(document, resource, embedded_fields)
        documents.append(document)

    response[app.config['ITEMS']] = documents
    return response, None


def getitem_internal(resource, **lookup):
    response = {}

    req = parse_request(resource)
    embedded_fields = resolve_embedded_fields(resource, req)
    # req.show_deleted = True
    document = app.data.find_one(resource, req, **lookup)
    if not document:
        abort(404)
    response = document
    build_response_document(document, resource, embedded_fields)

    for field in embedded_fields:
        embedded_document = document.get(field)
        if isinstance(embedded_document, dict):
            embedded_last_updated = last_updated(embedded_document)
            if embedded_last_updated > last_modified:
                last_modified = embedded_last_updated

    return response, None


from .auth import blueprint as _auth
from .swagger import blueprint as _swagger
from .schedules import blueprint as _schedules
from .students import blueprint as _students
from .calendar import blueprint as _calendar
from .tutor_stats import blueprint as _tutor_stats
from .forgot_password import blueprint as _forgot_password
from .progress import blueprint as _progress
