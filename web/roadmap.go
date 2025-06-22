package main

type Article struct {
	Title   string
	Content string
}

type Subtopic struct {
	Title       string
	Description string
	Keywords    []string
	Articles    []Article
}

type Topic struct {
	Title       string
	Description string
	Keywords    []string
	Subtopics   []Subtopic
}

type Roadmap struct {
	Area   string
	Topics []Topic
}

var roadmap = Roadmap{
	Area: "Backend-разработка",
	Topics: []Topic{
		{
			Title:       "Языки программирования и среды выполнения",
			Description: "Основные языки и платформы для разработки backend-сервисов.",
			Keywords:    []string{"Go", "Golang", "Java", "Python", "Py", "Node.js", "JS", "JavaScript", "JVM", "контейнеры"},
			Subtopics: []Subtopic{
				{
					Title: "Go", Description: "Компилируемый, производительный язык от Google.",
					Keywords: []string{"goroutines", "channels", "modules", "go modules", "Golang", "concurrency"},
					Articles: []Article{
						{
							Title:   "article title 1",
							Content: "article content 1 aaaa aaaaaa a a a dfsfkldsjaf;lkjsdf;lk jsadlk;j sdljf slkdjflksdjf sadf;lj sdlkjf lksadjf",
						},
						{
							Title:   "article title 2",
							Content: "article content 2 aaaa aaaaaa a a a dfsfkldsjaf;lkjsdf;lk jsadlk;j sdljf slkdjflksdjf sadf;lj sdlkjf lksadjf",
						},
						{
							Title:   "article title 3",
							Content: "article content 3 aaaa aaaaaa a a a dfsfkldsjaf;lkjsdf;lk jsadlk;j sdljf slkdjflksdjf sadf;lj sdlkjf lksadjf",
						},
					},
				},
				{
					Title: "Java", Description: "Широко используемый язык на JVM для корпоративных систем.",
					Keywords: []string{"Java", "Spring", "JVM", "JPA", "Maven", "Gradle", "Garbage Collection", "Multi-threading"},
				},
				{
					Title: "Python", Description: "Язык со множеством фреймворков для быстрого прототипирования.",
					Keywords: []string{"Python", "py", "pip", "Django", "Flask", "asyncio", "Производительность", "GIL"},
				},
				{
					Title: "Node.js", Description: "JavaScript-окружение для серверных приложений.",
					Keywords: []string{"Node.js", "js", "JavaScript", "npm", "Express", "event loop", "Асинхронность", "ECMAScript", "TypeScript"},
				},
			},
		},
		{
			Title:       "Веб-протоколы и API",
			Description: "Протоколы обмена данными между клиентом и сервером.",
			Keywords:    []string{"HTTP", "REST", "GraphQL", "gRPC", "WebSocket", "API"},
			Subtopics: []Subtopic{
				{
					Title: "REST", Description: "Архитектурный стиль для HTTP API.",
					Keywords: []string{"REST", "CRUD", "status codes", "JSON", "идентификаторы", "HATEOAS"},
				},
				{
					Title: "GraphQL", Description: "Язык запросов и среда исполнения от Facebook.",
					Keywords: []string{"GraphQL", "schema", "resolvers", "Apollo", "фрагменты", "интроспекция"},
				},
				{
					Title: "gRPC", Description: "RPC-фреймворк с HTTP/2 и Protobuf.",
					Keywords: []string{"gRPC", "Protobuf", "streaming", "IDL", "interop", "HTTP/2"},
				},
				{
					Title: "WebSocket", Description: "Двусторонняя связь по протоколу TCP.",
					Keywords: []string{"WebSocket", "реалтайм", "socket.io", "пинг/понг", "ws"},
				},
			},
		},
		{
			Title:       "Хранение данных",
			Description: "Базы данных и хранилища для backend-приложений.",
			Keywords:    []string{"SQL", "NoSQL", "ORM", "хранилище", "Database"},
			Subtopics: []Subtopic{
				{
					Title:       "PostgreSQL",
					Description: "Реляционная СУБД с расширенным функционалом.",
					Keywords:    []string{"PostgreSQL", "ACID", "indexes", "JSONB", "CTE", "MVCC"},
				},
				{
					Title:       "MongoDB",
					Description: "Документо-ориентированная NoSQL база данных.",
					Keywords:    []string{"MongoDB", "collections", "replication", "sharding", "агрегация", "NoSQL"},
				},
				{
					Title:       "Redis",
					Description: "In-memory key-value store для кэширования и очередей.",
					Keywords:    []string{"Redis", "кэш", "pub/sub", "persist", "in-memory"},
				},
				{
					Title:       "Elasticsearch",
					Description: "Поисковый движок для полнотекстового поиска.",
					Keywords:    []string{"Elasticsearch", "индексация", "shard", "репликация", "search"},
				},
			},
		},
		{
			Title:       "Контейнеризация и оркестрация",
			Description: "Развертывание и управление контейнерами в продакшне.",
			Keywords:    []string{"Docker", "Kubernetes", "контейнеры", "кластер", "containerization", "orchestration"},
			Subtopics: []Subtopic{
				{
					Title:       "Docker",
					Description: "Контейнеризация приложений.",
					Keywords:    []string{"Docker", "Dockerfile", "images", "volumes", "сети", "containers"},
				},
				{
					Title:       "Kubernetes",
					Description: "Оркестрация контейнеров в кластере.",
					Keywords:    []string{"Kubernetes", "Pods", "Deployments", "Ingress", "Helm", "service mesh"},
				},
				{
					Title:       "Helm",
					Description: "Менеджер пакетов для Kubernetes.",
					Keywords:    []string{"Helm", "charts", "templates", "releases", "package manager"},
				},
			},
		},
		{
			Title:       "CI/CD и DevOps",
			Description: "Автоматизация сборки, тестирования и деплоя.",
			Keywords:    []string{"CI", "CD", "Jenkins", "GitLab CI", "GitHub Actions", "DevOps"},
			Subtopics: []Subtopic{
				{
					Title:       "Jenkins/GitLab CI/GitHub Actions",
					Description: "Платформы для автоматизации пайплайнов.",
					Keywords:    []string{"github", "gitlab", "pipelines", "stages", "runners", "automation"},
				},
				{
					Title:       "Ansible/Terraform",
					Description: "Infrastructure as Code.",
					Keywords:    []string{"IaC", "playbooks", "модули", "provisioning", "Terraform", "Ansible"},
				},
			},
		},
		{
			Title:       "Сообщения и очереди",
			Description: "Асинхронная коммуникация между сервисами.",
			Keywords:    []string{"RabbitMQ", "Kafka", "очереди", "pub/sub", "messaging"},
			Subtopics: []Subtopic{
				{
					Title:       "RabbitMQ",
					Description: "Сообщения по AMQP.",
					Keywords:    []string{"RabbitMQ", "exchange", "queue", "binding", "routing"},
				},
				{
					Title:       "Apache Kafka",
					Description: "Распределенная стриминговая платформа.",
					Keywords:    []string{"Kafka", "topics", "brokers", "consumer groups", "stream processing"},
				},
			},
		},
		{
			Title:       "Мониторинг и логирование",
			Description: "Отслеживание состояния и логов приложений.",
			Keywords:    []string{"Prometheus", "Grafana", "ELK", "monitoring", "logging"},
			Subtopics: []Subtopic{
				{
					Title:       "Prometheus/Grafana",
					Description: "Сбор метрик и визуализация.",
					Keywords:    []string{"Prometheus", "metrics", "dashboards", "alertmanager", "visualization"},
				},
				{
					Title:       "ELK Stack",
					Description: "Лог-менеджмент и аналитика.",
					Keywords:    []string{"Elasticsearch", "Logstash", "Kibana", "beats", "logging"},
				},
			},
		},
		{
			Title:       "Архитектура и паттерны",
			Description: "Проектирование надежных и масштабируемых систем.",
			Keywords:    []string{"Monolith", "Microservices", "DDD", "Event-driven", "architecture"},
			Subtopics: []Subtopic{
				{
					Title:       "Микросервисы",
					Description: "Разделение приложения на независимые сервисы.",
					Keywords:    []string{"Microservices", "API-gateway", "service discovery", "circuit breaker", "resilience"},
				},
				{
					Title:       "Event-driven architecture",
					Description: "Реакция на события внутри системы.",
					Keywords:    []string{"Event-driven", "CDC", "event sourcing", "CQRS", "message-driven"},
				},
			},
		},
		{
			Title:       "Обеспечение отказоустойчивости и масштабируемости",
			Description: "Поддержка высокой доступности и производительности.",
			Keywords:    []string{"Load balancing", "High availability", "Auto-scaling", "resilience"},
			Subtopics: []Subtopic{
				{
					Title:       "Балансировка нагрузки",
					Description: "Распределение трафика между инстансами.",
					Keywords:    []string{"NGINX", "HAProxy", "round-robin", "load balancer"},
				},
				{
					Title:       "Auto-scaling/Failover",
					Description: "Автоматическое изменение числа инстансов.",
					Keywords:    []string{"auto-scaling", "horizontal scaling", "health checks", "failover"},
				},
			},
		},
		{
			Title:       "Безопасность и соответствие",
			Description: "Практики для защиты данных и сервисов.",
			Keywords:    []string{"OWASP", "GDPR", "security", "compliance", "safety"},
			Subtopics: []Subtopic{
				{
					Title:       "OAuth2/JWT",
					Description: "Аутентификация и авторизация.",
					Keywords:    []string{"OAuth2", "JWT", "tokens", "scopes", "refresh"},
				},
				{
					Title:       "Шифрование и секреты",
					Description: "TLS, KMS, Vault.",
					Keywords:    []string{"TLS", "SSL", "KMS", "Vault", "X.509", "AES", "Hashicorp Vault"},
				},
			},
		},
	},
}
