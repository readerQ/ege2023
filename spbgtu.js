javascript: (() => {
    //const documentHTML = document.documentElement.outerHTML;
    const matches = document.getElementsByClassName("enroll-lists-table__table");

    for (let item of matches){
        item.innerHTML = "<b>XXX</b>"
    }

 
    // const flatMatches = Array.from(matches).map((item) => item[0]);
    // const uniqueMatches = Array.from(new Set(flatMatches));
    // // if (uniqueMatches.length > 0) {
    //     const result = uniqueMatches.join(% 27\n % 27); alert(result);
    // } else { alert(% 27No emails found! % 27); }

})();