#!bin/bash
echo "......."
echo "Starting setup replica set database"
echo "......."

mongo1=`getent hosts ${MONGO1} | awk '{ print $1 }'`
mongo2=`getent hosts ${MONGO2} | awk '{ print $1 }'`
mongo3=`getent hosts ${MONGO3} | awk '{ print $1 }'`

echo "Waiting for startup.."
until mongo mongodb://${mongo1}:27017 --eval 'quit(db.runCommand({ ping: 1 }).ok ? 0 : 2)' &>/dev/null; do
  printf '.'
  sleep 1
done

echo "Started...."
mongo mongodb://${mongo1}:27017 <<EOF
    var cfg = {
        "_id": "ca",
        "members": [
            {
                "_id": 0,
                "host": "${mongo1}:27017"
            },
            {
                "_id": 1,
                "host": "${mongo2}:27017",
                "priority": 0
            },
            {
                "_id": 2,
                "host": "${mongo3}:27017",
                "priority": 0
            }
        ]
    };
    rs.initiate(cfg);
    rs.conf();
EOF