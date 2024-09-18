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
       // URLを指定して画面遷移
       window.location.href = url;
    });

});
