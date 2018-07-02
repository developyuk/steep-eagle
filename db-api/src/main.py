from eve import Eve
from eve_sqlalchemy import SQL
from eve_sqlalchemy.validation import ValidatorSQL

from tables import Base
from links import before_returning_items, before_returning_item
from auth import MyAuth, blueprint

app = Eve(auth=None, validator=ValidatorSQL, data=SQL)
app.register_blueprint(blueprint)

# bind SQLAlchemy
db = app.data.driver
Base.metadata.bind = db.engine
db.Model = Base

# app.on_fetched_resource += before_returning_items
# app.on_fetched_item += before_returning_item

if __name__ == '__main__':
  app.run(host='0.0.0.0', debug=True)
