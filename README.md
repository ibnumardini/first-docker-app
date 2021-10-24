# First Docker App 😃.

My process of learning docker.

## Cooking 🍳

Use this step, for build this app.

```bash
# build our app
docker build --tag=users-app:1.0.0-alpha .

# create container using the images we created earlier
docker container create --name=users-app --env-file .env -p 1234:8000 users-app:alpha1.0.0-alpha

# check our container
docker container ps
```

## Served 🚀

Use this step, for running/served this app.

```bash
# start docker container
docker container start users-app

# access this url
http://localhost:1234
```

## Learning Resources 📚
- [ProgrammerZamanNow (Eko Kurniawan Khannedy)](https://www.youtube.com/channel/UC14ZKB9XsDZbnHVmr4AmUpQ)
- [DuckDuckGo](hhttps://duckduckgo.com/?q=learn+docker)