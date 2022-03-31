var memoList;

document.addEventListener('DOMContentLoaded', function(){ 
    updateMemoList()

    insertBtn = document.getElementById("insert-memo-button")

    insertBtn.addEventListener("click", function () {
        content = document.getElementById("memo-content").value


        insertMemoRequest = new XMLHttpRequest()
        insertMemoRequest.onload = function () {
            updateMemoList()
        }
        insertMemoRequest.open("POST", "/api/memo/")
        insertMemoRequest.send(`{"content":"${content}"}`)

        document.getElementById("memo-content").value = ""
    })

}, false);

function updateMemoList () {
    getMemoRequest = new XMLHttpRequest();
    let tableBody = document.getElementById("memo-list-body")
    tableBody.innerHTML = ""

    getMemoRequest.onload = function () {
        console.log(this.responseText)

        memos = JSON.parse(this.responseText)
        memoList = memos
        displayMemos(memos)
    }

    getMemoRequest.open("GET", "/api/memo/");
    getMemoRequest.send();

}

function deleteMemo (id) {
    deleteMemoRequest = new XMLHttpRequest()

    deleteMemoRequest.onload = function () {
        updateMemoList()
    }

    deleteMemoRequest.open("DELETE", `/api/memo/${id}`)
    deleteMemoRequest.send()
}

function displayMemos (memos) {
    memos.map(memo => {
        let tableBody = document.getElementById("memo-list-body")

        row = document.createElement("tr")
        
        idCol = document.createElement("th")
        idCol.innerText = memo.id
        row.appendChild(idCol)

        uuidCol = document.createElement("td")
        uuidCol.innerText = memo.lastEditedBy
        row.appendChild(uuidCol)

        contentCol = document.createElement("td")
        contentCol.innerText = memo.content
        row.appendChild(contentCol)

        deleteCol = document.createElement("td")
        deleteButton = document.createElement("input")
        deleteButton.type = "button"
        deleteButton.value = "Delete"
        deleteButton.addEventListener("click", function () {
            //Delete
            memoId = memo.id
            deleteMemo(memoId)

        });


        row.appendChild(deleteButton)

        tableBody.appendChild(row)
    })
}