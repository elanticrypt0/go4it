# go4it

Es la base para crear diferentes aplicaciones que utilicen bases de datos y algunas cosita más.

Es mi SDK base para poder crear herramientas que no requieran tener una web online sino que se manejen por commandos.

## IMPORTANTE

go4it utiliza algunas librerías como DBmanager que se utilizan en gasonline. Tengo que encontrar la forma de poder organizar esos mismos proyectos como una base que pueda ser compartida y utilizada por estos proyectos. Por el momento será copiar y pegar.

# Dev mode

Por defecto go4it/dbmanager.go incluye base de datos mysql,postgress y sqlite. En caso de no querer alguna deberán borrarse para obtener un ejecutable más ligero.

# filepaths

Carpetas dentro de los archivos

./go4it: hay código común utilizado para la apps por ejemplo dbmanager para conectarse a las bases de datos o leer archivos json.
./src: archivos específicos del programa.
./cmd: comandos para enviar y ejecutar en la app. Las funciones dentro de src deberán ser llamadas por aquí o en el main directamente.

# Archivo de configuración

Una de las funciones importantes son los archivos de configuración. Estos archivos permiten tener varias bases de datos y conexiones en una misma app.

Para iniciar una app

    app := go4it.NewApp("")

Y para conectarse con una db. EL nombre de la conexión es el que tiene en el archivo toml [db.NOMBRE]

    app.Connect2Db("local")

Eso crea una nueva conexión a la que se accede por app.DB[ID] donde ID es el órden de en que hicimos la conexión. De toda formas si se quiere verificar que conexión es se puede retornar el nombre con

    app.DB[0].Name

Se pueden setear dos conexiones para que sean la primaria y secundaria con

    app.DB.SetPrimary(0)

Donde 0 es el índice de la conexión realizada.


# Usar los métodos de go4it

Si sólo se quiere extender alguna parte de go4it sólo hay que mirar las structs de db y las func de filehandler

# light build

al momento de compilar agregar estos dos parámetros para quitar accesorios de debug y que el programa sea más liviano:

    go build -ldflags "-w -s"

Compilar para diferentes plataformas:

    GOOS=windows GOARCH=386 go build -ldflags "-w -s"

Para ver las plataformas disponibles
    go tool dist list

Para ver la plataforma actual
    go env GOOS GOARCH