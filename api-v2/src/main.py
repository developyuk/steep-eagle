import os
from eve import Eve
import my_eve
from eve_swagger import swagger

app = Eve(
    # auth=my_eve.MyAuth,
    settings=os.path.abspath('settings.py'),
    data=my_eve.GoogleCloudstore
    )

app.register_blueprint(swagger)

if __name__ == '__main__':
    app.run(host='0.0.0.0', debug=os.environ['DEBUG'])
