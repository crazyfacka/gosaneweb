function displayBlob(data) {
    const urlCreator = window.URL || window.webkitURL;
    const imageUrl = urlCreator.createObjectURL(data);
    document.querySelector('#image_content').src = imageUrl;

    if (!!document.querySelector('.jcrop-stage')) {
        return;
    }

    const stage = Jcrop.attach('image_content');
    stage.listen('crop.change', (widget, e) => {
        let w = stage.el.clientWidth;
        let wMM = payload.features.X.values[1];

        let h = stage.el.clientHeight;
        let hMM = payload.features.Y.values[1];

        let posMM = {
            top: Math.round(widget.pos.y * hMM / h),
            left: Math.round(widget.pos.x * wMM / w),
            x: Math.round((widget.pos.x + widget.pos.w) * wMM / w),
            y: Math.round((widget.pos.y + widget.pos.h) * hMM / h),
        };

        document.querySelector('#x').value = posMM.x;
        document.querySelector('#y').value = posMM.y;
        document.querySelector('#top').value = posMM.top;
        document.querySelector('#left').value = posMM.left;
    });
}

function downloadBlob(data) {
    const urlCreator = window.URL || window.webkitURL;
    const imageUrl = urlCreator.createObjectURL(data);
    const a = document.createElement('a');

    a.href = imageUrl;
    a.download = 'scan.png';
    document.body.appendChild(a);

    a.click();
    a.remove();
}

function doScan(handler) {
    fetch('/scan', {
        method: 'POST',
        cache: 'no-cache',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(payload),
    }).then(response => response.blob())
    .then(handler)
    .catch((error) => {
        console.error('Error:', error);
    });
}

function getPreview() {
    const oldDPI = payload.features.RESOLUTION.to_use;
    const oldX = document.querySelector('#x').value;
    const oldY = document.querySelector('#y').value;
    const oldTop = document.querySelector('#top').value;
    const oldLeft = document.querySelector('#left').value;
    
    const mode = document.querySelector('#mode');

    payload.features.RESOLUTION.to_use = '75';
    payload.features.X.to_use = '215';
    payload.features.Y.to_use = '297';
    payload.features.T.to_use = '0';
    payload.features.L.to_use = '0';

    payload.features.MODE.to_use = mode.options[mode.selectedIndex].value;
    
    doScan(displayBlob);

    payload.features.RESOLUTION.to_use = oldDPI;
    payload.features.X.to_use = oldX;
    payload.features.Y.to_use = oldY;
    payload.features.T.to_use = oldTop;
    payload.features.L.to_use = oldLeft;
}

function scanFile() {
    const mode = document.querySelector('#mode');
    const dpi = document.querySelector('#resolution');
    const brightness = document.querySelector('#brightness').value;
    const contrast = document.querySelector('#contrast').value;
    const x = document.querySelector('#x').value;
    const y = document.querySelector('#y').value;
    const top = document.querySelector('#top').value;
    const left = document.querySelector('#left').value;

    payload.features.MODE.to_use = mode.options[mode.selectedIndex].value;
    payload.features.RESOLUTION.to_use = dpi.options[dpi.selectedIndex].value;
    payload.features.BRIGHTNESS.to_use = brightness;
    payload.features.CONTRAST.to_use = contrast;
    payload.features.X.to_use = x;
    payload.features.Y.to_use = y;
    payload.features.T.to_use = top;
    payload.features.L.to_use = left;

    doScan(downloadBlob);
}