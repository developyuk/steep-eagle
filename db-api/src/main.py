import os

from eve import Eve
from eve_sqlalchemy import SQL
from eve_sqlalchemy.validation import ValidatorSQL
from eve_swagger import swagger, add_documentation

from tables import Base
# from links import before_returning_items, before_returning_item
from hooks import modules, classes
# from mqtt import sessions
from auth import MyAuth
from blueprints import auth, schedules, students, tutor_stats, calendar

app = Eve(auth=MyAuth, validator=ValidatorSQL, data=SQL,
          settings=os.path.abspath('settings.py'))
app.register_blueprint(swagger)
app.register_blueprint(auth.blueprint)
app.register_blueprint(schedules.blueprint)
app.register_blueprint(students.blueprint)
app.register_blueprint(tutor_stats.blueprint)
app.register_blueprint(calendar.blueprint)

# bind SQLAlchemy
db = app.data.driver
Base.metadata.bind = db.engine
db.Model = Base

app.on_update_modules += modules.before_patch_item
app.on_insert_modules += modules.before_post_item
app.on_update_classes += classes.before_patch_item
app.on_inserted_classes += classes.after_post_item
# app.on_fetched_resource += before_returning_items
# app.on_fetched_item += before_returning_item
# app.on_inserted += sessions.sessions_tutor_students

add_documentation({
    'securityDefinitions': {
        "MyAuth": {"type": "apiKey", "name": "Authorization", "in": "header"}
    }
})

for resource, rd in app.config['DOMAIN'].items():
    if (rd.get('disable_documentation')
            or resource.endswith('_versions')):
        continue

    methods = rd['resource_methods']
    url = '/%s' % rd['url']
    for method in methods:
        add_documentation({'paths': {url: {method.lower(): {
            "security": [{"MyAuth": []}],
            'produces': ['application/json']
        }}}})

    methods = rd['item_methods']
    item_id = '%sId' % rd['item_title'].lower()
    url = '/%s/{%s}' % (rd['url'], item_id)
    for method in methods:
        add_documentation({'paths': {url: {method.lower(): {
            "security": [{"MyAuth": []}],
            'produces': ['application/json']
        }}}})

if __name__ == '__main__':
    app.run(host='0.0.0.0', debug=os.environ['DEBUG'])
