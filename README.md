# Demo Database connection

การทดลองนี้ใช้ SQLite เป็นตัวอย่าง จุดประสงค์เพื่อให้เห็นพฤติกรรมการทำงานกับ package database/sql เป็นสำคัญ
สามารถดูตัวอย่างเพิ่มเติมได้ที่ http://go-database-sql.org/

วิธีการเชื่อมต่อและการใช้ parameterize ต่างๆจะขึ้นอยู่กับ driver ของ database ที่เลือกใช้
เพราะฉะนั้น เวลาเลือกใช้ driver ใด ควรไปอ่านวิธีใช้ใน repository ของ driver นั้นๆก่อน

สามารถใช้ https://github.com/avelino/awesome-go#database-drivers
เป็นแหล่งอ้างอิงในการค้นหา driver ได้อีกช่องทาง
