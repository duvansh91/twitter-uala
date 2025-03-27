# Como instalarlo
Para levantar el proyecto en local se requiere tener instalado `docker` y `docker-compose`.
* Ubicado en la raíz del proyecto, Ejecutar el comando `docker-compose up`. Esto creará un contenedor para la API y un contenedor para la base de datos (mongoDB).

* Por defecto el API estara corriendo en el puerto 8080, puede cambiarse en el archivo `config/local.yml` en la propiedad `server_port`.

# Como usarlo

* Seguir a un usuario:
  * Path: `http://localhost:8080/twitter-uala/user/{user_id}/follow`
  * Método HTTP: `POST`
  * Body:  
  
  ```json 
      {
        "user_id": "id user to follow"
      }
  ```
    
    Posibles códigos de respuesta:
    
    * `200` OK -> Satisfactorio.
    
    * `400` Bad Request -> Cuando el formato del id en el body es inválido.

    * `500` Internal error -> Cuando hay algún error interno.

---

* Publicar un Tweet:
  * Path: `http://localhost:8080/twitter-uala/user/{user_id}/publish`
  * Método HTTP: `POST`
  * Body:  
  
  ```json 
      {
        "content": "content of the tweet"
      }
  ```
    
    Posibles códigos de respuesta:
    
    * `200` OK -> Satisfactorio.
    
    * `500` Internal error -> Cuando hay algún error interno.

---

* Ver Timeline:
  * Path: `http://localhost:8080/twitter-uala/user/{user_id}/timeline`
  * Método HTTP: `GET`
    
    Posibles códigos de respuesta:
    
    * `200` OK -> Satisfactorio.
    
    * `500` Internal error -> Cuando hay algún error interno.
