# Go_CRUD_API
This is an API build using ```golang```, ```gin``` to handle routing and ```mongodb``` as a database.

## #Installation Guide
  * Clone this repo into your local directory using following command
  ```bash
  git clone git@github.com:UtkarshM-hub/Go_CRUD_API.git
  ```
  * update the environment variable ```MONGODB_URI``` in ```.env``` file to your cluster URI. To setup the cluster checkout <a href="https://studio3t.com/knowledge-base/articles/mongodb-atlas-tutorial/" alt="go site">this</a> site.
  * If you haven't installed go yet then check the <a href="https://go.dev/doc/install" alt="go site">go site</a> to download go
  * Run following command
  ```bash
  go mod download && go run .
  ```
### Congratulations🥳🎊 you've completed the setup of GO_CRUD_API on your local machine

## #Routes
1. ```/```
 * ```/``` route creates a post in the data base.
 * Make a ```POST``` request to ```localhost:8080``` and the structue of request body should be like following
 ```JSON
 {
  "Title":"ENTER_THE_TITLE",
  "Article":"ENTER_THE_ARTICLE_DATA",
  }
  ```
2. ```/post/:postId```
 * This route sends the post as a response
 * The METHOD of request is ```GET```
3. ```/update/:postId```
  * This route updates the post
  * The METHOD of request is ```UPDATE```
  * The request body should be like one showed in ```/``` route
4. ```/delete/:postId```
 * This route is used to delete the post
 * The METHOD of request is ```DELETE```

## #Screenshots for example
<details>
<summary>Request</summary>
<img src="https://user-images.githubusercontent.com/70505181/190604332-d21c14d1-e4f7-4876-b2ab-c8e9dcaae7ef.png" alt="screenshot"/>
</details>
<details>
<summary>Database</summary>
<img src="https://user-images.githubusercontent.com/70505181/190604426-9bc045c3-125b-40ce-96db-79e09660372f.png" alt="screenshot"/>
</details>


