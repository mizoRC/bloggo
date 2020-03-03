db.createUser(
    {
        "user": "bloggo",
        "pwd": "bloggo",
        "roles": [
            {
                "role": "readWrite",
                "db": "bloggo"
            }
        ]
    }
)