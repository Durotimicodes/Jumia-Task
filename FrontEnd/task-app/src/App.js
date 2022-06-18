import './App.css';

export default function App() {
  return (
    <div className="App">
<div id="container">
  <header>
    <h1>JUMIA TASK</h1>
  </header>
  <div className="child-container">
    <section className="image">
      <img className="jumia-header-logo" src="/Jumia-logo - Copy.png" alt="jumia header logo" />
    </section>
    <section className="table-container">
      <h2>PHONE NUMBERS</h2>
      <div className="select-filters">
        <select name="Country" id="country" className="filter-country">
          <option value="Cameroon">Select a Country</option>
          <option value="Cameroon">Cameroon</option>
          <option value="Ethopia">Ethiopia</option>
          <option value="Morocco">Morocco</option>
          <option value="Mozambique">Mozambique</option>
          <option value="Uganda">Uganda</option>
        </select>
        <select name="Country" id="country" className="check-valid-num">
          <option value="Cameroon">Check Validity</option>
          <option value="Cameroon">Valid</option>
          <option value="Ethopia">Not Valid</option>
        </select>
      </div>
      <div className="phone-num-table">
        <table>
          <tbody><tr>
              <th>Country</th>
              <th>State</th>
              <th>Country code</th>
              <th>Phone number</th>
            </tr>
            <tr>
              <td>Alfreds Futterkiste</td>
              <td>Maria Anders</td>
              <td>Germany</td>
              <td>Germany</td>
            </tr>
            <tr>
              <td>Centro comercial Moctezuma</td>
              <td>Francisco Chang</td>
              <td>Mexico</td>
              <td>Germany</td>
            </tr>
            <tr>
              <td>Ernst Handel</td>
              <td>Roland Mendel</td>
              <td>Austria</td>
              <td>Germany</td>
            </tr>
            <tr>
              <td>Island Trading</td>
              <td>Helen Bennett</td>
              <td>UK</td>
              <td>Germany</td>
            </tr>
            <tr>
              <td>Laughing Bacchus Winecellars</td>
              <td>Yoshi Tannamuri</td>
              <td>Canada</td>
              <td>Germany</td>
            </tr>
            <tr>
              <td>Magazzini Alimentari Riuniti</td>
              <td>Giovanni Rovelli</td>
              <td>Italy</td>
              <td>Germany</td>
            </tr>
          </tbody></table>
      </div>
      <div className="pagnination">PAGINATION</div>
    </section>
  </div>
  <footer>© 2022 by Oluwadurotimi Edmond Fagbuyi</footer>
</div>

    </div>
  );
}

