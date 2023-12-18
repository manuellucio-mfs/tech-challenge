EMPAQUETADO DE CODIGO

1. Una vez descargado el codigo se debe de posicionar sobre el directorio cmd
    - cd cmd

2. Realizar el compilado del codigo 
    - set GOOS=linux
    - go build -o main main.go

3. Realizar el comprimido en formato zip del codigo necesario para la ejecucion de lambda en aws
    - %USERPROFILE%\Go\bin\build-lambda-zip.exe -output main.zip main

AWS CONFIGURATION

1. Crear buckets dentro de s3, ya que estos tienen nombres unicos se colocaron dentro de parameter store para recibirlo como parametro
    - pending-transaction-tets
    - processed-transaction-test

2. Crear la tabla "transactions-test" (este nombre puede variar de ambiente a ambiente) BD en dynamo donde colocaremos "idAccount" con id de particion

3. Crear topic SNS "balance-test" de tipo estandar (guardar ARN de topic), agregar a el o los suscriptores necesariios. A dico correo electronico llegara un email que debera confirmar para poder enviar cualquier informacion desde AWS SNS

4. Crear parametro, dentro de parameter store crear un parametro de nombre "tech-challenge" el cual contiene un json con la siguiente estructura:
    - {
        "s3Docs": {
            "pendingTransaction": "pending-transaction-test",
            "processedTransaction": "processed-transaction-test"
        },
        "tableName": "transactions-test",
        "snsArn": "arn:aws:sns:us-east-1:296164203735:balance-test",
        "pathFile": "MOLM123.csv"
    }
    NOTA: es posible cambiar los valores ajustados a los que se hayan creado en el nuevo ambiente

5. Crear lambda role agregando permisos para s3, dynamoDB, parameter store, sns, cloudWatch y lambda (para este ejercicio se agregaron permisos fullAccess pero es importante delimitar ua los que e utilizan unicamente por cuestiones de seguridad)

6. Crear Lambda con desencadenador de tipo API Gateway sin autentificacion para este proposito y subir codigo (ya sea desde archivo en maquina o desde s3 del archivo con extension .zip)
    - Realizar test para verificar que todo se ha configurado adecuadamente

7. Para probar desde el endpoint generado, ir a la opcion de configuracion y seleccionar desencadenadores
    - Puuede probarse directamente al dar clic en el endpoint que se muestra
    - Copiar el endpoint en alguna plataforma como postman y realizar la peticion