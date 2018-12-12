import os

from eve_swagger import swagger
from eve import Eve

import my

import hooks
import blueprints

app = Eve(auth=my.JwtAuth,
          media=my.GoogleMediaStorage)

app.register_blueprint(swagger)
app.register_blueprint(blueprints._swagger)
app.register_blueprint(blueprints._auth)
app.register_blueprint(blueprints._forgot_password)
app.register_blueprint(blueprints._schedules)
app.register_blueprint(blueprints._students)
app.register_blueprint(blueprints._tutor_stats)
app.register_blueprint(blueprints._import)

app.on_insert += hooks.resource.on_insert       # pylint: disable=no-member
app.on_inserted += hooks.resource.on_inserted   # pylint: disable=no-member

app.on_insert += hooks.classes.on_insert    # pylint: disable=no-member

app.on_insert += hooks.image_default.on_insert  # pylint: disable=no-member

app.on_inserted += hooks.mqtt.on_inserted           # pylint: disable=no-member
app.on_deleted_item += hooks.mqtt.on_deleted_item   # pylint: disable=no-member

if __name__ == '__main__':
    app.run(host='0.0.0.0', debug=os.environ['DEBUG'])
