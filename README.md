<h1>LEGIT</h1>
<p>Legit adalah Go Framework yang dikembangkan oleh <a href="https://codingers.id/" target="_blank">CODINGERS.ID</a> sebagai framework yang ditujukan khusus untuk pemula belajar bahasa pemrograman Go. Framework ini kami tujukan untuk komunitas IT kami, tetapi siapapun boleh menggunakan sebagai bahan belajar dan juga bisa digunakan untuk pengembangan website dan aplikasi di production.</p>
<p>Framework ini jauh dari kata sempurna, jadi kami akan selalu dan terus berinovasi agar bisa menjadi framework yang berguna dan bermanfaat di dunia pemrograman.</p>

<hr>
<h2>Dikembangkan Dengan</h2>
<p>Framework ini dibangun dengan:</p>
<ul>
    <li><a href="https://go.dev/" target="_blank">Go</a></li>
    <li><a href="https://docs.gofiber.io/" target="_blank">Fiber</a></li>
    <li><a href="https://gorm.io/docs/index.html" target="_blank">GORM</a></li>
</ul>

<hr>
<h2>Prasarat</h2>
<p>Untuk bisa menggunakan framework ini, ada beberapa hal yang harus anda install terlebih dahulu.</p>
<ul>
    <li><a href="https://git-scm.com/downloads" target="_blank">GIT</a></li>
    <li><a href="https://go.dev/dl/" target="_blank">Bahasa Go versi >= 1.17</a></li>
    <li>Salah Satu Database Ini:
        <ul>
            <li><a href="https://www.mysql.com/downloads/" target="_blank">MySQL</a></li>
            <li><a href="https://mariadb.org/download" target="_blank">MariaDB</a></li>
            <li><a href="https://www.postgresql.org/download/" target="_blank">PostgreSQL</a></li>
        </ul></li>
    <li>Salah Satu Text Editor Ini:
        <ul>
            <li><a href="https://www.sublimetext.com/" target="_blank">SublimeText</a></li>
            <li><a href="https://code.visualstudio.com/" target="_blank">Visual Studio Code</a></li>
            <li><a href="https://cursor.sh/" target="_blank">Cursor</a></li>
        </ul></li>
    <li>Pendukung lainnya seperti Tools Manajemen Database, Browser dll</li>
</ul>

<hr>
<h2>Installasi</h2>
<p>Berikut ini langkah-langkah installasi legit framework:</p>
<ol>
    <li>Buka terminal/cmd/git bash.<pre><code>git clone https://github.com/codingersid/legit.git namaAplikasiAnda</code></pre></li>
    <li>Buka folder namaAplikasiAnda di Text Editor Anda</li>
    <li>Copy file <code>.env.example</code> kemudian ubah menjadi <code>.env</code></li>
    <li>Buatlah 1 database dan konfigurasi pada file <code>.env</code></li>
    <li>Buka terminal dan jalankan server.<pre><code>go run legit.go dev</code></pre></li>
    <li>Setelah berhasil, akses URL <code>http://127.0.0.1:3000/</code> di Browser Anda</li>
    <li>Selamat Legit Framework Berhasil Diinstall dan Dijalankan!</li>
</ol>

<hr>
<h2>Legit CLI</h2>
<p>Berikut ini adalah beberapa perintah yang bisa Anda jalankan di terminal dalam pengembangan projek dengan Legit:</p>
<ol>
    <li>Perintah untuk <code>help</code>.<pre><code>go run legit.go -h</code></pre></li>
    <li>Perintah untuk <code>menjalankan server</code>.<pre><code>go run legit.go dev</code></pre></li>
    <li>Perintah untuk cek <code>versi</code>.<pre><code>go run legit.go versi</code></pre></li>
    <li>Perintah untuk membuat <code>controller</code>.<pre><code>go run legit.go controller NamaControllerNya</code></pre></li>
    <li>Perintah untuk membuat <code>model</code>.<pre><code>go run legit.go model NamaModelNya</code></pre></li>
    <li>Perintah untuk membuat <code>request</code>.<pre><code>go run legit.go request NamaRequestNya</code></pre></li>
    <li>Perintah untuk membuat <code>middleware</code>.<pre><code>go run legit.go middleware NamaMiddlewareNya</code></pre></li>
    <li>Perintah untuk membuat <code>migration</code>.<pre><code>go run legit.go migration NamaTabelNya</code></pre></li>
    <li>Perintah untuk membuat <code>seeder</code>.<pre><code>go run legit.go seeder NamaSeederNya</code></pre></li>
</ol>

<hr>
<h2>Menjalankan Command legit di Terminal</h2>
<p>Apabila Anda merasa command dengan <code>go run legit.go [command]</code> terlalu panjang, Anda bisa mengatur di terminal agar bisa menjadi singkat menjadi <code>legit [command]</code></p>

<h3>Pengaturan di Mac/Unix/Linux</h3>
<ol>
    <li>Buka project Anda dengan Text Editor, lalu jalankan terminal.</li>
    <li>Jalankan perintah berikut ini:
        <pre><code>go install</code></pre>
        <pre><code>export PATH=$PATH:$(go env GOPATH)/bin</code></pre>
        <pre><code>source ~/.zshrc atau ~/.bashrc</code></pre>
    </li>
</ol>

<h3>Pengaturan di Windows</h3>
<ol>
    <li>Buka project Anda dengan Text Editor, lalu jalankan terminal.</li>
    <li>Jalankan perintah berikut ini:
        <pre><code>go install</code></pre>
    </li>
    <li>Tambahkan <code>C:\Go\bin</code> pada PATH Windows Anda.<a href="https://wahyu-ehs.medium.com/cara-install-golang-di-windows-5060aa2383a9">Menginstall Path Go di Windows</a></li>
</ol>

<hr>
<h2>Template Engine</h2>
<p>Untuk mempermudah pembuatan project, khususnya website, maka diperlukan template enginge. Legit menggunakan template engine dari <a href="https://docs.gofiber.io/template/django/">Django</a> yang disediakan oleh Fiber.</p>

<hr>
<h2>Kontribusi</h2>
<p>Apabila Ingin berkontribusi dalam pengembangan framework ini, silahkan anda Fork repositori ini.</p>

<hr>
<h2>Kontak</h2>
<p>Kontak kami ada di link pada akhir file ini, bisa hubungi kami melalui Instagram.</p>

<hr>
<h2>Support dan Sponsorship</h2>
<p>Apabila Anda ingin mensupport kami dalam bentuk finansial ataupun lainnya, kami terbuka untuk hal tersebut.</p>

<hr>
<h3 align="center">FOLLOW AKUN KAMI</h3>
<p align="center">
<a href="https://www.instagram.com/codingers.id/" target="_blank" rel="noopener noreferrer">INSTAGRAM</a>
&nbsp;|&nbsp;<a href="https://www.facebook.com/codingers.id" target="_blank" rel="noopener noreferrer">FACEBOOK</a>
&nbsp;|&nbsp;<a href="https://codingers.id/" target="_blank" rel="noopener noreferrer">WEBSITE</a>
&nbsp;|&nbsp;<a href="https://github.com/codingersid/" target="_blank" rel="noopener noreferrer">GITHUB</a>
</p>
