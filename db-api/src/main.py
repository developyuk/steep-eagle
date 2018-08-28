import os

from eve import Eve
from eve_sqlalchemy import SQL
from eve_swagger import swagger

from tables import Base
from hooks import classes, image_empty
# from mqtt import sessions
from customs import MyAuth, MyMediaStorage, MyValidator
from blueprints import auth, schedules, students, tutor_stats, calendar, swagger as my_swagger

app = Eve(auth=MyAuth, media=MyMediaStorage, validator=MyValidator, data=SQL,
          settings=os.path.abspath('settings.py'))
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

app.on_update_classes += classes.before_patch_item
app.on_inserted_classes += classes.after_post_item
app.on_fetched_resource += image_empty.on_fetched_resource
app.on_fetched_item += image_empty.on_fetched_item
# app.on_inserted += sessions.sessions_tutor_students

if __name__ == '__main__':
    app.run(host='0.0.0.0', debug=os.environ['DEBUG'])
