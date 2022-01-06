class Table extends React.Component {
    constructor(props) {
        super(props)
        this.state = {
            customers: null,
            availableCountries: [],
            availableStatus: [],
            filters: {
                country_name: '',
                status: ''
            }
        }
        this.getCustomers()
    }

    getCustomers(params = '') {
        fetch('http://localhost:8080/customers/search' + params)
            .then(response => response.json())
            .then(data => {
                this.setState({
                    customers: data.customers,
                    availableCountries: [...new Set(data.customers.map(item => item.country_name))],
                    availableStatus: [...new Set(data.customers.map(item => item.status))],
                    filters: this.state.filters
                })
            });
    }

    filterByCountry = (event) => {
        let searchParam
        let country = event.target.value

        this.setState({
            customers: this.state.customers,
            availableCountries: this.state.availableCountries,
            availableStatus: this.state.availableStatus,
            filters: {
                country_name: country,
                status: this.state.filters.status
            }
        })

        if (country === '') {
            searchParam = this.state.filters.status === '' ? '' : '?status=' + this.state.filters.status
        } else {
            let statusFilter = this.state.filters.status
            searchParam = statusFilter === '' ? '?country_name=' + country : '?status=' + statusFilter + '&country_name=' + country
        }

        this.getCustomers(searchParam)
    }

    filterByStatus = (event) => {
        let searchParam
        let status = event.target.value

        this.setState({
            customers: this.state.customers,
            availableCountries: this.state.availableCountries,
            availableStatus: this.state.availableStatus,
            filters: {
                country_name: this.state.filters.country_name,
                status: status
            }
        })

        if (status === '') {
            searchParam = this.state.filters.country_name === '' ? '' : '?country_name=' + this.state.filters.country_name
        } else {
            let countryNameFilter = this.state.filters.country_name
            searchParam = countryNameFilter === '' ? '?status=' + status : '?status=' + status + '&country_name=' + countryNameFilter
        }

        this.getCustomers(searchParam)
    }

    render() {
        return (
            this.state.customers ?
                <div className="container">
                    <div className="row">
                        <div className="col-xs-8 col-xs-offset-2 jumbotron text-center">
                            <h1>Customers</h1>
                            <div>
                                <div>
                                    Status
                                    <select onChange={this.filterByStatus}>
                                        <option></option>
                                        {this.state.availableStatus.map(item => {
                                            return (
                                                <option>{item}</option>
                                            );
                                        })}
                                    </select>
                                </div><div>
                                Country
                                <select onChange={this.filterByCountry}>
                                    <option></option>
                                    {this.state.availableCountries.map(item => {
                                        return (
                                            <option>{item}</option>
                                        );
                                    })}
                                </select>
                            </div>
                            </div>
                            <div className="table-responsive">
                                <table class="table">
                                    <thead>
                                    <tr>
                                        <th>Customer Name</th>
                                        <th>Country Name</th>
                                        <th>Country Code</th>
                                        <th>Phone Number</th>
                                        <th>Status</th>
                                    </tr>
                                    </thead>
                                    <tbody>
                                    {this.state.customers.map(item => {
                                        return (
                                            <tr key={item.phone_number}>
                                                <td>{item.customer_name}</td>
                                                <td>{item.country_name}</td>
                                                <td>{item.country_code}</td>
                                                <td>{item.phone_number}</td>
                                                <td>{item.status}</td>
                                            </tr>
                                        );
                                    })}
                                    </tbody>
                                </table>
                            </div>
                        </div>
                    </div>
                </div>
                : <p>Loading</p>
        );
    }
}

ReactDOM.render(<Table/>, document.getElementById("app"));