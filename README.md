# StocksApp

API en Go (Fiber) para la aplicación StocksApp, con conexión a CockroachDB y endpoints REST.

1. Clona el repositorio
2. Configura las variables de entorno (ver abajo)
3. Ejecuta:

#### A. Configurar variables de entorno tanto en backend como en frontend
Organiza el archivo .env.example en .env tanto en la carpeta de backend como en la de frontend. Recuerda que debes tener una cuenta en CockroachDB o usar una local.

```bash
cp .env.example .env
```

#### B. Alimenta la base de datos con data de la API de finnhub
- Primero verifica la conexion a la base de datos es correcta
 ```bash
 cd backend
go run cmd/testapi/main.go
```
Si todo esta bien, ejecuta:

```bash
cd backend
go run cmd/setupdb/main.go
```
Y espera que termine de ejecutar

#### C. Ejecutar el backend

- Para ejecutar con frontend en desarrollo

```bash
cd backend
go run cmd/api/main.go
cd ../frontend
npm install
npm run dev
```

- Para ejecutar backend con frontend estaticos

```bash
cd frontend
npm install
npm run build
cd ../backend
go run cmd/api/main.go
```