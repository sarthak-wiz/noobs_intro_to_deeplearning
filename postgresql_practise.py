import psycopg2
import os
import sys
from dotenv import load_dotenv

load_dotenv()

database = sys.argv[1]
user = os.environ.get("PGUSER")
password = os.environ.get("PGPASSWORD")
host = os.environ.get("PGHOST")
port = os.environ.get("PGPORT")


conn = None
conn = psycopg2.connect(database = database, user = user, password = password, host = host, port = port)
curr = conn.cursor()

query="select p.name, t.name, p.jersey_no from players p, teams t where p.team_id = t.team_id order by p.name desc, t.name desc"
curr.execute(query)
players = curr.fetchall()

print(players)

curr.close  
conn.close