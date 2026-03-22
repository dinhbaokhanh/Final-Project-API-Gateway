# PTIT Gateway - Nền Tảng Mạng Xã Hội Học Tập PTIT

Dự án xây dựng nền tảng mạng xã hội. Repository này chứa mã nguồn phần **API Gateway** cho hệ thống microservices.

## 1. Mục Tiêu Dự Án

*   **Mục đích:** Viết một API Gateway cho hệ thống microservices nhằm mục đích học tập và rèn luyện kỹ năng cũng như sử dụng cho đồ án môn học.

---

## 2. Luồng Hoạt Động Chi Tiết (Core Workflows)

### 2.1. Luồng Hỏi Đáp & Thảo Luận (Q&A Workflow)
1.  **Tạo bài:** Người dùng đăng câu hỏi, bài chia sẻ hoặc báo lỗi (Bug report). Có thể chọn đăng ẩn danh để giảm rào cản tâm lý ngại ngùng. Tích hợp trực tiếp Code Editor tại khung soạn thảo nếu cần hỏi về source code.
2.  **Phân loại (Auto-Tagging):** Bài đăng đi qua bộ vi xử lý AI để tự động nhận diện ngữ nghĩa, gán tag theo môn học, và phân loại bài viết (Hỏi đáp, Thảo luận, Chia sẻ).
3.  **Tương tác:** Cộng đồng tham gia upvote/downvote, comment trả lời, đánh giá dòng code (review comment).
4.  **Giải quyết:** Người đặt câu hỏi đánh dấu câu trả lời đúng/hữu ích nhất để vinh danh người hỗ trợ. Trạng thái bài đăng chuyển sang "Đã giải quyết" (Resolved).
5.  **Tóm tắt (AI Knowledge Extraction):** Định kỳ, hệ thống sẽ tự động tổng hợp, tóm tắt bài thảo luận thành các mục FAQ cấu trúc rõ ràng đưa vào kho tri thức chung.

### 2.2. Luồng Tìm Kiếm Nội Dung (Semantic Search Workflow)
1.  **Truy vấn:** Sinh viên nhập từ khóa, hệ thống đưa ra các gợi ý (Autocomplete).
2.  **Lọc đa chiều:** Áp dụng lọc kết hợp (Operators) theo các tiêu chí: Tag môn học, Loại bài, Trạng thái (đang thảo luận/đã giải quyết), Khoảng thời gian cụ thể.
3.  **Xử lý ngữ nghĩa:** Hệ thống sử dụng Semantic Search (Vector DB) phân tích truy vấn, tìm kiếm sâu trong tiêu đề, thẻ nội dung, nội dung file đính kèm, source code.
4.  **Trả kết quả:** Biểu thị danh sách các bài viết/nhóm liên quan, có thể sắp xếp theo hot/mới nhất/nhiều view nhất. Tính năng theo dõi Trending topics theo thời gian thực (tuần/tháng).

### 2.3. Luồng Thành Lập & Quản Lý Đồ Án Nhóm (Study Group Workflow)
1.  **Tìm bạn đồng hành (Partner Matching):** Hệ thống duyệt qua hồ sơ cá nhân, lịch rảnh, kỹ năng và lịch sử tương tác (bài đã upvote) của sinh viên để đề xuất những người cùng "tần số".
2.  **Tạo tổ chức:** Sinh viên tạo Nhóm công khai/riêng tư theo môn hoặc dự án cụ thể. Gửi yêu cầu/lời mời tham gia tới các thành viên tiềm năng.
3.  **Thiết lập quy tắc:** Phân quyền quản trị (Admin nhóm), thiết lập ghim bài nội bộ, nội quy nhóm.
4.  **Quản lý tiến độ (Task Board - Kanban):** Các thành viên cùng tạo Task, gán thành viên chịu trách nhiệm, đặt Deadline và di chuyển trạng thái task (To-do -> In Progress -> Done).
5.  **Lưu trữ & Lịch nhóm:** Upload tài nguyên dùng chung lên Cloud. Đồng bộ các deadline học tập với lịch (Google Calendar) giúp theo dõi chặt chẽ tiến trình làm đồ án.

### 2.4. Luồng Tích Lũy Điểm & Xếp Hạng (Gamification Workflow)
1.  **Ghi nhận thành tích:** Hệ thống lắng nghe các hành động tiêu biểu như: câu trả lời được chủ thread đánh dấu đúng, bài định hướng môn học được nhiều view/upvote, hoặc chia sẻ tài liệu tốt.
2.  **Quy đổi điểm (Reputation):** Cộng cho sinh viên mức điểm uy tín tương đương dựa trên độ hữu ích.
3.  **Vinh danh chuyên môn:** Cấp tự động các danh hiệu (Ví dụ: "Helper", "Code Reviewer", "Active Learner") hỗ trợ hiển thị nổi bật trên Leaderboard chung để kích thích thi đua học tập.

---

## 3. Tính Năng & Chức Năng Cốt Lõi

*   **Post & Nội dung:** Quản lý bài viết đa dạng trạng thái, Code Editor + Execution Sandbox (review diff), bộ lọc nội dung ẩn danh (Toxic detection).
*   **Quản lý Nhóm/Cộng đồng:** Task board, Lịch nhóm, chia sẻ tiến độ thiết kế Figma/Code nội bộ.
*   **Search & Đề xuất (AI tích hợp):** Graph matching algorithms (Neo4j), Semantic Search, Transformer xử lý tổng hợp tự động hỏi đáp.
*   **Quản lý người dùng & file:** Xác thực, phân quyền Dashboard kiểm duyệt chuyên sâu, Cloud storage tài liệu học tập cá nhân và tập thể.

---

## 4. Định Hướng Phát Triển API Gateway (Kiến trúc Configuration-Driven theo KrakenD)

Thay vì lập trình cứng (hard-code) các logic định tuyến và xử lý ngay trong mã nguồn gốc, API Gateway của hệ thống sẽ được tái cấu trúc và phát triển dựa trên triết lý cốt lõi của **KrakenD** (Configuration-driven API Gateway). Hướng đi này giúp tách biệt hoàn toàn phần xử lý logic (Engine) khỏi phần khai báo dịch vụ (Declaration), đồng thời hỗ trợ mạnh mẽ các mẫu kiến trúc microservices nâng cao.

Các tính năng cốt lõi sẽ tập trung xây dựng bao gồm:

### 4.1. Khởi chạy và Định tuyến bằng File Cấu Hình (Configuration-Driven Routing)
*   **Mô tả:** Thay vì code cứng (hard-code) các route như `"/api/users"` hay `"/api/orders"` vào trong bộ định tuyến (router) của Go, toàn bộ hệ thống sẽ sử dụng một file cấu hình duy nhất (ví dụ: `gateway.json` hoặc `gateway.yaml`).
*   **Hoạt động:** Khi Gateway khởi động, nó sẽ nạp (parse) file cấu hình này lên memory, tự động tạo các dynamic routes (đường dẫn động) và cấu hình các downstream services (dịch vụ đích) tương ứng.
*   **Lợi ích:** Dễ dàng thêm, sửa, xóa, hoặc quy hoạch lại API mapping mà không cần recompile mã nguồn.

### 4.2. Kiến trúc Backend for Frontend (BFF) thông qua Cấu Hình
*   **Mô tả:** Tạo ra các backend riêng biệt được may đo (tailor-made) cho từng loại frontend cụ thể (như Mobile App, Web App).
*   **Hoạt động:** Thay vì viết thêm Service trung gian, Gateway đảm nhận vai trò BFF. Cấu hình các endpoint trả về các payload khác nhau cho cùng một tài nguyên tùy thuộc vào client.

### 4.3. Cung cấp dữ liệu tập trung (Endpoint Aggregation / Data Fetching)
*   **Mô tả:** Hệ thống hỗ trợ nhận 1 request từ Client và tự động phát tách (fan-out) ra nhiều request gọi tới các Microservices.
*   **Hoạt động:** Gateway sử dụng Goroutines để gọi song song các service, trộn (merge) phản hồi lại thành một JSON object hợp nhất và trả về Client. Giảm độ trễ mạng và over-fetching.

### 4.4. Quản lý Middleware (Interceptor) tự động
*   **Mô tả:** Áp dụng hệ thống Middleware (Xác thực, CORS, Rate-Limiting, Logging) ở tầng Gateway.
*   **Hoạt động:** Tắt/bật Middleware thông qua cấu hình `gateway.json` ở cấp độ Toàn cục (Global) hoặc Cục bộ (Endpoint-level).

### 4.5. Kiến trúc Phi trạng thái (Stateless Architecture)
*   **Mô tả:** API Gateway không kết nối trực tiếp với Database hay Session Cache.
*   **Hoạt động:** Mọi thao tác ủy quyền được xác minh tính hợp lệ độc lập (verify JWT) hoặc forward về Auth Service, giúp Gateway dễ dàng scale.

---

## 5. Cấu Trúc Kỹ Thuật (Đang triển khai)
Dự án được viết bằng **Go**, hiện đang có cấu trúc như sau (sẽ được tái cấu trúc dần để đáp ứng kiến trúc mới):

### Cấu Trúc Thư Mục
```text
ptit-gateway/
|-- cmd/
|   `-- gateway/
|       `-- main.go          # Điểm khởi động
|-- internal/
|   |-- app/                 # Khởi tạo và chạy HTTP server
|   |-- config/              # Đọc biến môi trường
|   |-- routing/             # Khai báo route proxy -> backend services
|   |-- middleware/          # CORS, logger, recover, (sau này có rate-limit)
|   `-- proxy/               # Reverse proxy implementation
|-- go.mod
`-- README.md
```

### Các Biến Môi Trường (Mẫu)
- `PORT`: Cổng gateway (mặc định `8080` hoặc có thể thay đổi)
- `BACKEND_USERS`: URL service Users (Xác thực, Thông tin)
- `BACKEND_POSTS`: URL service Q&A, Thảo luận
- `BACKEND_GROUPS`: URL service Nhóm học tập, Kanban
- `BACKEND_SEARCH`: URL service Semantic Search & Analytics

### Build & Run
```bash
go mod tidy
go run ./cmd/gateway
```
