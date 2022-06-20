import { useEffect, useState } from "react";
import axios from "axios";
import _ from "lodash";

const pageSize = 5;

export default function Posts() {

    const [posts, setPosts] = useState([]);

    const [checkValue, setCheckValue] = useState("");

    const [checkCountry, setCheckCountry] = useState("");

    const [paginatedPost, setpaginatedPost] = useState("");

    const [currentPage, setCurrentPage] = useState(1)

    let y = checkCountry
    let x = checkValue

   const url = `http://localhost:8081/api/v1/user/mobile/${y}/${x}`

    useEffect(()=>{
        axios.get(url)

        .then(response=>{
            console.log(response.data.data, "Response")

            setPosts(response.data.data)

            setpaginatedPost(_(response.data.data).slice(0).take(pageSize).value())
        })
        .catch(
            error=>{
                console.log(error)
            }
        )
    },[checkValue, checkCountry])
    
    const pageCount = posts? Math.ceil(posts.length/pageSize) : 0;

    if (pageCount === 1) return null;

    const pages  = _.range(1, pageCount+1)

    const pagination = (pageNo)=>{
        setCurrentPage(pageNo);

        const startIndex = (pageNo - 1) * pageSize;

        const paginatedPost = _(posts).slice(startIndex).take(pageSize).value();

        setpaginatedPost(paginatedPost)
    }

  return (

    <div style={{width:"70vw", margin:"auto", border:"2px solid blue"}}>

        <h1 style={{textAlign:"center", fontSize:"55px"}}>JUMIA TASK</h1>

        <div className="select-filters">
    
        <select name="Country" id="country" className="filter-country"
                value={checkCountry} onChange={(e) => setCheckCountry(e.target.value)}>
          <option value="">Select Country</option>

          <option value="237">Cameroon</option>

          <option value="251">Ethopia</option>

          <option value="212">Morocco</option>

          <option value="258">Mozambique</option>

          <option value="256">Uganda</option>

        </select>

        
        <select name="Country" id="country" className="check-valid-num" style={{marginLeft:"50px"}}

        value={checkValue} onChange={(e) => setCheckValue(e.target.value)}>
            
          <option>Select Status</option>

          <option>OK</option>

          <option>NOK</option>

        </select>

      </div>
        
            <table className="table">

                <thead>

                    <tr>

                        <th>Country</th>

                        <th>State</th>

                        <th>Country code</th>

                        <th>Phone number</th>

                        <th>Status</th>

                    </tr>

                </thead>

                <tbody>
                    
                    { paginatedPost.length === 0 ?  (<h1>Select the dropdown list</h1>):(paginatedPost.map((post, index)=>{
                        return(
                            <tr key={index}>

                            <td>{post.country}</td>

                            <td>{post.state}</td>

                            <td>{post.countryCode}</td>

                            <td>{post.mobileNumber}</td>

                            <td>
                                <p className={post.state  === "OK" ? "btn btn-success": "btn btn-danger"}>
                                    {post.state === "OK" ? "Valid": "Not Valid"}
                                </p>

                            </td>

                            </tr>
                    )})
                    )}
                </tbody>

            </table>
        <nav className="d-flex justify-content-center">

            <ul className="pagination">
                {
                    pages.map((page)=>(
                    <li className={
                        page === currentPage ? "page-item active":"page-item"
                    }>
                        <p className="page-link" onClick={()=>pagination(page)}>
                        {page}
                        </p>
                        
                        </li>
                    ))
                }
            </ul>
        </nav>
        <footer>© 2022 by Oluwadurotimi Edmond Fagbuyi</footer>

        <section className="image">

      <img className="jumia-header-logo" src="/Jumia-logo - Copy.png" alt="jumia header logo" style={{width:"100%"}}/>

    </section>
        
    </div>
  )
}
