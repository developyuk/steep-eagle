import os

from eve import Eve
from eve_sqlalchemy import SQL
from eve_swagger import swagger

from tables import Base
from customs import MyAuth, MyMediaStorage, MyValidator
import hooks
import blueprints

app = Eve(auth=MyAuth, media=MyMediaStorage, validator=MyValidator,
          data=SQL, settings=os.path.abspath('settings.py'))

app.register_blueprint(swagger)
app.register_blueprint(blueprints.bp_swagger)
app.register_blueprint(blueprints.bp_auth)
app.register_blueprint(blueprints.bp_schedules)
app.register_blueprint(blueprints.bp_students)
app.register_blueprint(blueprints.bp_tutor_stats)
app.register_blueprint(blueprints.bp_calendar)

db = app.data.driver
Base.metadata.bind = db.engine
db.Model = Base

app.on_update_classes += hooks.class_students.on_update             # pylint: disable=no-member
app.on_insert_classes += hooks.class_students.on_insert             # pylint: disable=no-member
app.on_inserted_classes += hooks.class_students.on_inserted         # pylint: disable=no-member

app.on_fetched_resource += hooks.image_default.on_fetched_resource  # pylint: disable=no-member
app.on_fetched_item += hooks.image_default.on_fetched_item          # pylint: disable=no-member

app.on_inserted += hooks.mqtt.on_inserted                           # pylint: disable=no-member
app.on_deleted_item += hooks.mqtt.on_deleted_item                   # pylint: disable=no-member

app.on_insert += hooks.tutor_attend.on_insert                       # pylint: disable=no-member
app.on_deleted_item += hooks.tutor_attend.on_deleted_item           # pylint: disable=no-member

app.on_inserted += hooks.notifications.on_inserted                  # pylint: disable=no-member
app.on_deleted_item += hooks.notifications.on_deleted_item          # pylint: disable=no-member

app.on_insert += hooks.guardians.on_insert                          # pylint: disable=no-member
app.on_inserted += hooks.guardians.on_inserted                      # pylint: disable=no-member
app.on_update += hooks.guardians.on_update                          # pylint: disable=no-member

if __name__ == '__main__':
    app.run(host='0.0.0.0', debug=os.environ['DEBUG'])
