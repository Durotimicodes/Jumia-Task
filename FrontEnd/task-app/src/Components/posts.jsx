import { useEffect, useState } from "react";
import axios from "axios";




export default function Posts() {

    const [posts, setPosts] = useState([]);
    const [checkValue, setCheckValue] = useState("");
    const [checkCountry, setCheckCountry] = useState("");
    let y = checkCountry
    let x = checkValue
   console.log(x,"This is X")
   console.log(y,"This is Y")

   const url = `http://localhost:8081/api/v1/user/mobile/${y}/${x}`
   console.log(url, "THIS IS THE URL")

    useEffect(()=>{
        axios.get(url)
        .then(response=>{
            console.log(response.data.data, "Response")
            setPosts(response.data.data)
        })
        .catch(
            error=>{
                console.log(error)
            }
        )
    },[checkValue, checkCountry])
    
  return (
    <div>
        <div className="select-filters">
        <select name="Country" id="country" className="filter-country"
        value={checkCountry} onChange={(e) => setCheckCountry(e.target.value)}
        >
          <option value="">Select Country by Code</option>
          <option value="237">Cameroon</option>
          <option value="251">Ethopia</option>
          <option value="212">Morocco</option>
          <option value="258">Mozambique</option>
          <option value="256">Uganda</option>
        </select>

        <select name="Country" id="country" className="check-valid-num"
        value={checkValue} onChange={(e) => setCheckValue(e.target.value)}>
            
          <option>Select Validity</option>
          <option>OK</option>
          <option>NOK</option>
        </select>
      </div>
        
            <table>
                <thead>
                    <tr>
                        <th>Country</th>
                        <th>State</th>
                        <th>Country code</th>
                        <th>Phone number</th>
                    </tr>
                </thead>
                <tbody>
                    
                    { posts.length === 0 ?  (<h1>Select the dropdown list</h1>):(posts.map((post, index)=>{
                        return(
                            <tr key={index}>
                            <td>{post.country}</td>
                            <td>{post.state}</td>
                            <td>{post.countryCode}</td>
                            <td>{post.mobileNumber}</td>
                            </tr>
                    )})
                    )}
                </tbody>
            </table>
    
    </div>
  )
}
