from eve import Eve
from eve_sqlalchemy import SQL
from eve_sqlalchemy.validation import ValidatorSQL
import os

from tables import Base
# from links import before_returning_items, before_returning_item
from mqtt import sessions
from auth import MyAuth
from blueprints import auth, schedules, students, tutor_stats

app = Eve(auth=MyAuth, validator=ValidatorSQL, data=SQL, settings=os.path.abspath('settings.py'))
app.register_blueprint(auth.blueprint)
app.register_blueprint(schedules.blueprint)
app.register_blueprint(students.blueprint)
app.register_blueprint(tutor_stats.blueprint)

# bind SQLAlchemy
db = app.data.driver
Base.metadata.bind = db.engine
db.Model = Base

# app.on_fetched_resource += before_returning_items
# app.on_fetched_item += before_returning_item
# app.on_inserted += sessions.sessions_tutor_students

if __name__ == '__main__':
  app.run(host='0.0.0.0', debug=os.environ['DEBUG'])
