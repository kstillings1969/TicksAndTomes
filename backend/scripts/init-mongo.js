// Initialize MongoDB with application collections and indexes
db.createCollection("players");
db.players.createIndex({ "user_id": 1 }, { unique: true });
db.players.createIndex({ "email": 1 }, { unique: true });
db.players.createIndex({ "empire.clan_id": 1 });
db.players.createIndex({ "empire.status": 1 });
db.players.createIndex({ "empire.created_at": 1 });
db.players.createIndex({ "stars": 1 });

db.createCollection("clans");
db.clans.createIndex({ "name": 1 }, { unique: true });
db.clans.createIndex({ "created_by": 1 });
db.clans.createIndex({ "members": 1 });

db.createCollection("market_listings");
db.market_listings.createIndex({ "seller_id": 1 });
db.market_listings.createIndex({ "resource_type": 1 });
db.market_listings.createIndex({ "created_at": 1 });
db.market_listings.createIndex({ "active": 1 });

db.createCollection("chat_messages");
db.chat_messages.createIndex({ "clan_id": 1 });
db.chat_messages.createIndex({ "sender_id": 1 });
db.chat_messages.createIndex({ "created_at": 1 });

db.createCollection("attack_history");
db.attack_history.createIndex({ "defender_id": 1 });
db.attack_history.createIndex({ "attacker_id": 1 });
db.attack_history.createIndex({ "timestamp": 1 });

print("Database initialization complete!");
