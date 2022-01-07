class Table extends React.Component {
    constructor(props) {
        super(props)
        this.state = {
            customers: null,
            limit: 10,
            offset: 0,
            total: 0,
            filters: {
                country_name: '',
                status: ''
            }
        }
        this.getCustomers(this.state.limit, this.state.offset, this.state.filters)
    }

    getCustomers(limit = this.state.limit, offset = this.state.offset, params) {
        let baseUrl = 'http://localhost:8080/customers/search'
        let searchParams = ''

        if (params.status !== '') {
            searchParams += '&status=' + params.status
        }

        if (params.country_name !== '') {
            searchParams += '&country_name=' + params.country_name
        }

        let fullUrl = `${baseUrl}?limit=${limit}&offset=${offset}${searchParams}`

        fetch(fullUrl)
            .then(response => response.json())
            .then(data => {
                this.setState({
                    customers: data.customers,
                    limit: data.limit,
                    offset: data.offset,
                    total: data.total,
                    filters: this.state.filters
                })
                console.log(data.total)
            });
    }

    filterByCountry = (event) => {
        let country = event.target.value
        let params = {
            country_name: country,
            status: this.state.filters.status
        }

        this.setState({
            customers: this.state.customers,
            limit: this.state.limit,
            offset: 0,
            total: this.state.total,
            filters: params
        })

        this.getCustomers(this.state.limit, 0, params)
    }

    filterByStatus = (event) => {
        let status = event.target.value
        let params = {
            country_name: this.state.filters.country_name,
            status: status
        }

        this.setState({
            customers: this.state.customers,
            limit: this.state.limit,
            offset: 0,
            total: this.state.total,
            filters: params
        })

        this.getCustomers(this.state.limit, 0, params)
    }

    previousPage = () => {
        let previousPage = this.state.offset - this.state.limit

        if (previousPage >= 0) {
            this.setState({
                customers: this.state.customers,
                limit: this.state.limit,
                offset: previousPage,
                total: this.state.total,
                filters: this.state.filters
            })
            this.getCustomers(this.state.limit, previousPage, this.state.filters)
        }
    }

    nextPage = () => {
        let nextPage = this.state.offset + this.state.limit

        if (nextPage < this.state.total) {
            this.setState({
                customers: this.state.customers,
                limit: this.state.limit,
                offset: nextPage,
                total: this.state.total,
                filters: this.state.filters
            })
            this.getCustomers(this.state.limit, nextPage, this.state.filters)
        }
    }

    changeLimit = (event) => {
        let newLimit = event.target.value

        this.setState({
            customers: this.state.customers,
            limit: newLimit,
            offset: this.state.offset,
            total: this.state.total,
            filters: this.state.filters
        })

        this.getCustomers(newLimit, this.state.offset, this.state.filters)
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
                                        <option>Valid</option>
                                        <option>Invalid</option>
                                    </select>
                                </div>
                                <div>
                                    Country
                                    <select onChange={this.filterByCountry}>
                                        <option></option>
                                        <option>Cameroon</option>
                                        <option>Ethiopia</option>
                                        <option>Morocco</option>
                                        <option>Mozambique</option>
                                        <option>Uganda</option>
                                    </select>
                                </div>
                                <div>
                                    Display size
                                    <select onChange={this.changeLimit}>
                                        <option>10</option>
                                        <option>20</option>
                                        <option>40</option>
                                        <option>100</option>
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
                                    {this.state.customers.map((item) => {
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
                            <div>
                                <div>
                                    <button onClick={this.previousPage}>
                                        Previous Page
                                    </button>
                                </div>
                                <div>
                                    <button onClick={this.nextPage}>
                                        Next Page
                                    </button>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
                : <p>Loading</p>
        );
    }
}

ReactDOM.render(<Table/>, document.getElementById("app"));