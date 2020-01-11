
var booksAPI = '/books'

var searchForm = document.querySelector('.search')
var searchInput = document.querySelector('.search-field')
var bookLists = document.querySelector('.books-list')

const state = {}

class Book {
    constructor(ISBNumber, name, writer, translator, publisher, edition_note, print_year, no_of_page, updated, img = '') {
        this.ISBNumber = ISBNumber
        this.name = name
        this.writer = writer
        this.translator = translator
        this.publisher = publisher
        this.edition_note = edition_note
        this.print_year = print_year
        this.no_of_page = no_of_page
        this.updated = updated
        this.img = img
    }
}

const clearBookLists = () => {
    bookLists.innerHTML = ``
}

const clearSearchInput = () => {
    searchInput.value = ``
}

const clearLoader = () => {
    const loader = document.querySelector('.loader')
    if (loader) loader.parentElement.removeChild(loader)
}

const renderLoader = (el) => {
    const loader = `<div class="loader"></div>`
    el.insertAdjacentHTML('afterBegin', loader)
}

const renderBookLists = (books) => {
    books.map((book, index) => {
        let markup = `
                    <div class="book" id="${index}">
                            <img class="book-img" src="${book.img}">
                            <div class="book-detail">
                                <h4>${book.name}</h4>
                                <p>ผู้เขียน: ${book.writer}</p>
                                ${book.translator !== '' ? `<p>ผู้แปล: ${book.translator}</p>` : ``}
                                <p>ผู้จัดจำหน่าย: ${book.publisher}</p>
                                <p>ปีที่พิมพ์: ${book.print_year}</p>
                                <p>จำนวนหน้า: ${book.no_of_page}</p>
                                <p>Updated: ${book.updated}</p>
                            </div>
                        </div>`
        bookLists.insertAdjacentHTML('afterBegin',markup)
    })
}

const ctrlSearch = async () => {
    const searchVal = (searchInput.value).toLowerCase()
    console.log(searchVal)
    const query = searchVal
    if (query !== '') {
        try {
            clearBookLists()
            renderLoader(bookLists)
            clearSearchInput()
            state.book = new Book(book)
            await state.poke.getPokeAllDetail()
            if (state.poke.name) {
                clearLoader(bookLists)
                renderBookLists(state.book)
            }
        } catch (error) {
            console.log(`ctrlSearch ${error}`)
        }
    }
}

const ctrlStater = async () => {
    try {
        const books = await fetch(booksAPI).then(books => books.json())
        console.log(books)
        books.map((book) => {
            return new Book(book.ISBNumber, book.name, book.writer, book.translator, book.publisher, book.edition_note, book.print_year, book.no_of_page, book.updated)
        })
        renderBookLists(books)
    } catch (error) {
        console.log(`getBookLists ${error}`)
    }
}

const init = () => {
    searchForm.addEventListener('submit', function (event) {
        event.preventDefault()
        ctrlSearch()
    })
    ctrlStater()
}

init()