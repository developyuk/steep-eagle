import os

from eve_swagger import swagger
from eve import Eve

import my
import hooks
import blueprints

app = Eve(auth=my.JwtAuth,
          settings=os.path.abspath('settings.py'),
          data=my.GoogleCloudstore,
          media=my.GoogleMediaStorage)

app.register_blueprint(swagger)
app.register_blueprint(blueprints._auth)
app.register_blueprint(blueprints._students)
app.register_blueprint(blueprints._schedules)
app.register_blueprint(blueprints._forgot_password)
app.register_blueprint(blueprints._import)
app.register_blueprint(blueprints._swagger)

app.on_fetched_resource += hooks.resource.on_fetched_resource       # pylint: disable=no-member
app.on_fetched_item += hooks.resource.on_fetched_item               # pylint: disable=no-member
app.on_insert += hooks.resource.on_insert                           # pylint: disable=no-member

app.on_fetched_resource += hooks.image_default.on_fetched_resource  # pylint: disable=no-member
app.on_fetched_item += hooks.image_default.on_fetched_item          # pylint: disable=no-member

if __name__ == '__main__':
    app.run(host='0.0.0.0', debug=os.environ['DEBUG'])
