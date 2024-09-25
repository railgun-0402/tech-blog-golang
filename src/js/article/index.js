'use strict'

// DOM Treeの構築が完了次第処理開始
document.addEventListener('DOMContentLoaded', () => {
    const deleteBtns = document.querySelectorAll('.articles__item-delete');

    const csrfToken = document.getElementsByName('csrf')[0].content;

    const deleteArticle = id => {
        let statusCode;

        fetch(`/${id}`, {
            method: 'DELETE',
            headers: {'X-CSRF-Token': csrfToken}
        })
            .then(res => {
                statusCode = res.status;
                return res.json();
            })
            .then(data => {
                console.log(JSON.stringify(data));
                if (statusCode === 200) {
                    document.querySelector(`.articles__item-${id}`).remove();
                }
            })
            .catch(err => console.error(err));
    };

    // それぞれの削除ボタンに対し、イベントリスナーを設定
    for (let elm of deleteBtns) {
        elm.addEventListener('click', event => {
            event.preventDefault();

            // 削除ボタンのカスタムデータ属性からIDを取得して引数に渡す
            deleteArticle(elm.dataset.id);
        })
    }
})
