'use strict'

// DOM Treeの構築が完了次第処理開始
document.addEventListener('DOMContentLoaded', () => {
    const deleteBtns = document.querySelectorAll('.articles__item-delete');
    const moreBtn = document.querySelector('.page__more');
    const articles = document.querySelector('.articles');
    const articleTmpl = document.querySelector('.articles__item-tmpl').firstElementChild;

    const csrfToken = document.getElementsByName('csrf')[0].content;

    const deleteArticle = id => {
        let statusCode;

        fetch(`/api/articles/${id}`, {
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

    // もっとみるボタンにイベントリスナーを設定します。
    moreBtn.addEventListener('click', event => {
        event.preventDefault();

        // もっとみるボタンのカスタムデータ属性からカーソルの値を取得します。
        const cursor = moreBtn.dataset.cursor;

        // カーソルが取得できない場合や 0 以下の数値だった場合は、
        // もっとみるボタンを画面から削除して処理を終了します。
        if (!cursor || cursor <= 0) {
            moreBtn.remove();
            return;
        }

        // Fetch API を利用して非同期リクエストを実行します。
        let statusCode;
        fetch(`/api/articles?cursor=${cursor}`)
            .then(res => {
                statusCode = res.status;
                return res.json();
            })
            .then(data => {
                console.log(JSON.stringify(data));
                // リクエストに成功し、記事一覧データが配列で返ってきた場合
                if (statusCode == 200 && Array.isArray(data)) {
                    // 表示する記事がこれ以上存在しない場合は、
                    // もっとみるボタンを画面から削除して処理を終了します。
                    if (data.length == 0) {
                        moreBtn.remove();
                        return;
                    }

                    // 記事の HTML 要素をまとめるためのフラグメントを作成します。（記事のリスト）
                    const fragment = document.createDocumentFragment();

                    // 記事一覧データをループで処理します。
                    data.forEach(article => {
                        // 個々の記事の HTML 要素を格納するフラグメントを作成します。（個別記事）
                        const frag = document.createDocumentFragment();

                        // 記事の HTML 要素のテンプレートからクローンを作成し、
                        // フラグメントの子要素として追加します。
                        frag.appendChild(articleTmpl.cloneNode(true));

                        // 記事の各 HTML 要素に対して、クラス・属性値・テキストを設定します。
                        frag.querySelector('article').classList.add(`articles__item-${article.id}`);
                        frag.querySelector('.articles__item').setAttribute('href', `/articles/${article.id}`);
                        frag.querySelector('.articles__item-title').textContent = article.title;
                        frag.querySelector('.articles__item-date').textContent = article.created.split('T')[0]; //+年-月-日のみを抽出

                        // デリートボタンに対して、カスタムデータ属性やイベントリスナーを設定します。
                        const deleteBtnElm = frag.querySelector('.articles__item-delete');
                        deleteBtnElm.dataset.id = article.id;
                        deleteBtnElm.addEventListener('click', event => {
                            event.preventDefault();
                            deleteArticle(article.id);
                        });

                        // 記事リストのフラグメントの子要素として個別記事のフラグメントを追加します。
                        fragment.appendChild(frag);
                    });

                    // もっとみるボタンのカスタムデータ属性に設定してあるカーソルの値を更新します。
                    moreBtn.dataset.cursor = data[data.length - 1].id;

                    // 記事一覧の HTML 要素の子要素に記事リストのフラグメントを追加して画面に表示します。
                    articles.appendChild(fragment);
                }
            })
            .catch(err => console.error(err));
    });
})
