'use strict';

// DOM TREEの構築が完了したら処理開始
document.addEventListener('DOMContentLoaded', () => {
    // DOM APIを利用してHTML要素を取得する
    const inputs = document.getElementsByTagName('input');
    const form = document.forms.namedItem('article-form');
    const saveBtn = document.querySelector('.article-form__save')
    const cancelBtn = document.querySelector('.article-form__cancel');
    const previewOpenBtn = document.querySelector('.article-form__open-preview');
    const previewCloseBtn = document.querySelector('.article-form__close-preview');
    const articleFormBody = document.querySelector('.article-form__body');
    const articleFormPreview = document.querySelector('.article-form__preview');
    const articleFormBodyTextArea = document.querySelector('.article-form__input--body');
    const articleFormPreviewTextArea = document.querySelector('.article-form__preview-body-contents');
    const errors = document.querySelector('.article-form__errors');
    const errorTmpl = document.querySelector('.article-form__error-tmpl').firstElementChild;

    // 新規作成画面か編集画面をURLから判定
    const mode = { method: '', url: '' };
    if (window.location.pathname.endsWith('new')) {
        // 新規作成時のHTTPメソッドはPOST
        mode.method = 'POST';
        mode.url = '/';
    } else if (window.location.pathname.endsWith('edit')) {
        // 更新時にHTTPメソッドはPATCHを利用
        mode.method = 'PATCH';
        mode.url = `/${window.location.pathname.split('/')[1]}`;
    }
    const { method, url } = mode;
    // CSRFトークンを取得
    const csrfToken = document.getElementsByName('csrf')[0].content;

    // input要素にフォーカスが合った状態でEnterが押されるとformが送信される
    for (let elm of inputs) {
        elm.addEventListener('keydown', event => {
            if (event.keyCode && event.keyCode === 13) {
                // キーが押された際のデフォルトの挙動をキャンセル
                event.preventDefault();

                return false;
            }
        });
    }

    // 保存処理実行イベント
    saveBtn.addEventListener('click', event => {
       event.preventDefault();

       // 前回のバリデーションエラー表示を削除
        errors.innerHTML = null;

       const formData = new FormData(form);
       let status;

       fetch(url, {
           method: method,
           headers: {'X-CSRF-Token': csrfToken },
           body: formData
       })
           .then(res => {
               status = res.status;
               return res.json();
           })
           .then(body => {
               console.log(JSON.stringify(body));

               if (status === 200) {
                   window.location.href = url;
               }

               if (body.ValidationErrors) {
                   showErrors(body.ValidationErrors);
               }
           })
           .catch(err => console.error(err));
    });

    // バリデーションエラー表示
    const showErrors = messages => {
        if (Array.isArray(messages) && messages.length !== 0) {
            const fragment = document.createDocumentFragment();

            messages.forEach(message => {
                const frag = document.createDocumentFragment();

                // テンプレートをクローンしてフラグメントに追加
                frag.appendChild(errorTmpl.cloneNode(true))

                frag.querySelector('.article-form__error').innerHTML = message;

                fragment.appendChild(frag);
            });

            // エラーメッセージの表示エリアにメッセージを追加
            errors.appendChild(fragment);
        }
    }

    // プレビューを開くイベントを設定
    previewOpenBtn.addEventListener('click', event => {
        // form本文に入力されたMarkdownをHTMLに変換して、プレビューに埋め込む
        // articleFormPreviewTextArea.innerHTML = articleFormBodyTextArea.value;
        articleFormPreviewTextArea.innerHTML = md.render(articleFormBodyTextArea.value);
        articleFormBody.style.display = 'none'; // 入力フォーム非表示
        articleFormPreview.style.display = 'grid'; // プレビュー表示
    });

    // プレビュー閉じるイベントを設定
    previewCloseBtn.addEventListener('click', event => {
       articleFormBody.style.display = 'grid'; // 入力フォーム表示
       articleFormPreview.style.display = 'none'; // プレビュー非表示
    });

    // 前ページに戻るイベント
    cancelBtn.addEventListener('click', event => {
       event.preventDefault();

       console.log(url);
       // URLを指定して画面遷移
       window.location.href = url;
    });

});
