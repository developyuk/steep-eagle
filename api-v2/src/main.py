import os
from eve import Eve
import my_eve
from eve_swagger import swagger
import blueprints

app = Eve(auth=my_eve.MyAuth,
          settings=os.path.abspath('settings.py'),
          data=my_eve.GoogleCloudstore)

app.register_blueprint(swagger)
app.register_blueprint(blueprints._auth)

if __name__ == '__main__':
    app.run(host='0.0.0.0', debug=os.environ['DEBUG'])
