import os

from eve import Eve
from eve_sqlalchemy import SQL
from eve_swagger import swagger

from tables import Base
from hooks import class_students, image_default, mqtt
# from mqtt import sessions
from customs import MyAuth, MyMediaStorage, MyValidator
from blueprints import auth, schedules, students, tutor_stats, calendar, swagger as my_swagger

app = Eve(auth=MyAuth, media=MyMediaStorage, validator=MyValidator,
          data=SQL, settings=os.path.abspath('settings.py'))

app.register_blueprint(swagger)
app.register_blueprint(my_swagger.blueprint)
app.register_blueprint(auth.blueprint)
app.register_blueprint(schedules.blueprint)
app.register_blueprint(students.blueprint)
app.register_blueprint(tutor_stats.blueprint)
app.register_blueprint(calendar.blueprint)

# bind SQLAlchemy
db = app.data.driver
Base.metadata.bind = db.engine
db.Model = Base

app.on_update_classes += class_students.before_patch_item
app.on_insert_classes += class_students.before_post_item
app.on_inserted_classes += class_students.after_post_item
app.on_fetched_resource += image_default.on_fetched_resource
app.on_fetched_item += image_default.on_fetched_item
app.on_inserted += mqtt.on_inserted
app.on_deleted_item += mqtt.on_deleted_item

if __name__ == '__main__':
    app.run(host='0.0.0.0', debug=os.environ['DEBUG'])
