print('Start #################################################################');

const username = _getEnv('DB_USERNAME')
const password = _getEnv('DB_PASSWORD')
const dbname = _getEnv('DB_NAME')
const collection = _getEnv('DB_COLLECTION')

db = db.getSiblingDB(dbname);
db.createUser(
  {
    user: username,
    pwd: password,
    roles: [{ role: 'readWrite', db: dbname }]
  },
);
db.createCollection(collection);
print('END #################################################################');
