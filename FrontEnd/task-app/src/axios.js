import axios from "axios";

//Next we create an instance of it 
const instance = axios.create({
    //here we have our configurations

    baseURL: 'http://localhost:8081/api/v1'
})

export default instance

